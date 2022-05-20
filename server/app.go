package server

import (
	"fmt"
	"net/http"
	"time"

	"net"

	"github.com/astronautsid/astro-ims-be/infrastructures"
	"github.com/astronautsid/astro-ims-be/interfaces"
	"github.com/astronautsid/astro-ims-be/repositories"
	"github.com/astronautsid/astro-ims-be/resources"
	"github.com/astronautsid/astro-ims-be/services"
	"google.golang.org/grpc"

	"github.com/astronautsid/astro-ims-be/controllers"
	//healthpb "github.com/astronautsid/astro-proto/golang/pb/healthcheck"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	InventoryService interfaces.InterfaceInventoryService
	server           *grpc.Server
}

//GracefulStop gracefully stop GRPC server
func (s *GRPCServer) GracefulStop() {
	s.server.GracefulStop()
}

func (s *GRPCServer) Serve(port string) error {
	tcpListener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	return s.server.Serve(tcpListener)
}

var (
	severityCheckerFunc = func(code codes.Code) zapcore.Level {
		if code == codes.OK {
			return zapcore.InfoLevel
		}
		return zapcore.ErrorLevel
	}

	grpc_metrics = grpc_prometheus.NewServerMetrics()
	zapLogger, _ = zap.NewProduction()
)

//Run for running service (initializes both GRPC and HTTP server)
func Run(conf *resources.AppConfig, dbObj map[string]interfaces.IDatabase, redisdb map[string]interfaces.IRedis) (*GRPCServer, *http.Server) {
	// Shared options for the logger, with a custom gRPC code to log level function.
	grpcLoggerOptions := []grpc_zap.Option{
		grpc_zap.WithLevels(severityCheckerFunc),
	}
	// Make sure that log statements internal to gRPC library are logged using the logrus Logger as well.
	grpc_zap.ReplaceGrpcLogger(zapLogger)

	grpcServer := GRPCServer{
		server: grpc.NewServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_prometheus.StreamServerInterceptor,
			)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_validator.UnaryServerInterceptor(),
				grpc_prometheus.UnaryServerInterceptor,
				grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
				grpc_zap.UnaryServerInterceptor(zapLogger, grpcLoggerOptions...),
			)),
		),
	}

	//healthpb.RegisterHealthServer(grpcServer.server, grpcServer)
	reflection.Register(grpcServer.server)

	// Initialize all metrics.
	grpc_metrics.EnableHandlingTimeHistogram()
	grpc_prometheus.EnableHandlingTimeHistogram()

	grpc_prometheus.Register(grpcServer.server)
	grpc_metrics.InitializeMetrics(grpcServer.server)

	httpServer := &http.Server{
		Addr:         ":" + conf.Core.Inventory.Port,
		Handler:      Router().Routing(conf, dbObj, redisdb),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return &grpcServer, httpServer
}

//ServiceInject is for injecting dependencies into each layer of service (from top to bottom : controller -> service -> repository)
func ServiceInject(
	config *resources.AppConfig,
	dbObj map[string]interfaces.IDatabase,
	redisdb map[string]interfaces.IRedis,
) *controllers.InventoryController {

	//new context logging should be initialized on the controller level
	//ctx := context.Background()
	webhookURL := config.Core.Inventory.Slack.WebhookURL
	webhookChannel := fmt.Sprintf("#%s", config.Core.Inventory.Slack.WebhookChannel)
	env := config.Core.Inventory.Environment

	slackConf := &infrastructures.SlackWebhook{
		SlackWebhookEnv:     env,
		SlackWebhookURL:     webhookURL,
		SlackWebhookChannel: webhookChannel,
	}

	var iSlackWebhook interfaces.ISlackWebhook
	iSlackWebhook = slackConf

	panicHandler := &infrastructures.PanicHandlerController{
		Slack:       iSlackWebhook,
		SlackConfig: slackConf,
	}

	var iPanicHandler interfaces.IPanicHandler
	iPanicHandler = panicHandler

	iHttp := infrastructures.HTTPCall{
		Conf: &config.HTTPConfig,
	}

	// iJwt := &infrastructures.ConfigJwt{
	// 	Ctx: ctx,
	// 	Env: env,
	// }

	inventoryRepository := &repositories.InventoryRepository{
		DB:    dbObj[resources.DatabasePostgreSQL],
		Redis: redisdb[resources.RedisDefault],
		HTTP:  &iHttp,
		Conf:  config,
	}

	inventoryService := &services.InventoryService{
		Repo: inventoryRepository,
		Conf: config,
		HTTP: &iHttp,
		//JWT:   iJwt,
	}

	inventoryController := controllers.InventoryController{
		Services:       inventoryService,
		Configurations: config,
		Slack:          iSlackWebhook,
		SlackConfig:    slackConf,
		PanicHandler:   iPanicHandler,
	}

	return &inventoryController
}
