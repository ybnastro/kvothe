package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SurgicalSteel/kvothe/infrastructures"
	"github.com/SurgicalSteel/kvothe/interfaces"
	"github.com/SurgicalSteel/kvothe/resources"
	"github.com/SurgicalSteel/kvothe/server"

	"github.com/go-redis/redis"
)

var (
	timeWaiting = (5 * time.Second)
)

func main() {
	conf := loadConfig()

	dbList := make(map[string]interfaces.IDatabase)
	dbPostgres := infrastructures.PostgreSQLHandler{}
	dbPostgres.ConnectDB(&conf.Core.Kvothe.DBPostgres.READ, &conf.Core.Kvothe.DBPostgres.WRITE)

	dbList[resources.DatabasePostgreSQL] = &dbPostgres

	redisdb := infrastructures.RedisHandler{
		Client: &redis.Client{},
	}
	redisdb.ConnectRedis(&conf.Core.Kvothe.Redis)

	redisList := make(map[string]interfaces.IRedis)
	redisList[resources.RedisDefault] = &redisdb

	httpServer := server.Run(conf, dbList, redisList)

	log.Printf("%s starting on HTTP Port %s and GRPC Port %s", conf.Core.Kvothe.Name, conf.Core.Kvothe.Port, conf.Core.Kvothe.GRPC.Port)
	ShutdownApp(httpServer, conf.Core.Kvothe.Name, conf.Core.Kvothe.GRPC.Port, dbList, redisList[resources.RedisDefault])

}

//ShutdownApp for shutting down server gracefully
func ShutdownApp(httpServer *http.Server, serverName, grpcPort string, db map[string]interfaces.IDatabase, redisdb interfaces.IRedis) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[ERROR][httpServer] ListenAndServe: %s\n", err)
		}
	}()

	<-done
	log.Printf("Stopping %s\n", serverName)
	ctx, cancel := context.WithTimeout(context.Background(), timeWaiting)
	defer func() {
		// extra handling here
		for _, v := range db {
			v.Close()
		}
		redisdb.Close()
		cancel()
	}()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to stop http service %s %+v\n", serverName, err)
	}

	log.Printf("%s stopped successfully\n", serverName)
}
