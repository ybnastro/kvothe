package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SurgicalSteel/kvothe/infrastructures"
	"github.com/SurgicalSteel/kvothe/interfaces"
	"github.com/SurgicalSteel/kvothe/repositories"
	"github.com/SurgicalSteel/kvothe/resources"
	"github.com/SurgicalSteel/kvothe/services"

	"github.com/SurgicalSteel/kvothe/controllers"
)

//Run for running service (initializes both GRPC and HTTP server)
func Run(conf *resources.AppConfig, dbObj map[string]interfaces.IDatabase, redisdb map[string]interfaces.IRedis) *http.Server {

	httpServer := &http.Server{
		Addr:         ":" + conf.Core.Kvothe.Port,
		Handler:      Router().Routing(conf, dbObj, redisdb),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return httpServer
}

//ServiceInject is for injecting dependencies into each layer of service (from top to bottom : controller -> service -> repository)
func ServiceInject(
	config *resources.AppConfig,
	dbObj map[string]interfaces.IDatabase,
	redisdb map[string]interfaces.IRedis,
) *controllers.KvotheController {

	//new context logging should be initialized on the controller level
	//ctx := context.Background()
	webhookURL := config.Core.Kvothe.Slack.WebhookURL
	webhookChannel := fmt.Sprintf("#%s", config.Core.Kvothe.Slack.WebhookChannel)
	env := config.Core.Kvothe.Environment

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

	kvotheRepository := &repositories.KvotheRepository{
		DB:    dbObj[resources.DatabasePostgreSQL],
		Redis: redisdb[resources.RedisDefault],
		HTTP:  &iHttp,
		Conf:  config,
	}

	kvotheService := &services.KvotheService{
		Repo: kvotheRepository,
		Conf: config,
		HTTP: &iHttp,
		//JWT:   iJwt,
	}

	kvotheController := controllers.KvotheController{
		Services:       kvotheService,
		Configurations: config,
		Slack:          iSlackWebhook,
		SlackConfig:    slackConf,
		PanicHandler:   iPanicHandler,
	}

	return &kvotheController
}
