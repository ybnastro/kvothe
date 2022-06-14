package main

import (
	"log"
	"os"

	"github.com/SurgicalSteel/kvothe/resources"
	"github.com/SurgicalSteel/kvothe/utils"

	"github.com/alexsasharegan/dotenv"
)

func loadConfig() *resources.AppConfig {
	err := dotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	var conf resources.AppConfig

	if err != nil {
		log.Fatal(err)
	}

	//GIN
	conf.GINMode = os.Getenv("GIN_MODE")

	//HTTP Config
	conf.HTTPConfig.Timeout = utils.GetInt(os.Getenv("HTTP_TIMEOUT_SECOND"))
	conf.HTTPConfig.DisableKeepAlive = utils.GetBool(os.Getenv("HTTP_DISABLE_KEEP_ALIVE"))

	conf.Core.Kvothe.Redis.URL = os.Getenv("KVOTHE_SERVICE_REDIS_URL")
	conf.Core.Kvothe.Redis.Port = utils.GetInt(os.Getenv("KVOTHE_SERVICE_REDIS_PORT"))
	conf.Core.Kvothe.Redis.DB = utils.GetInt(os.Getenv("KVOTHE_SERVICE_REDIS_DB"))
	conf.Core.Kvothe.Redis.Password = os.Getenv("KVOTHE_SERVICE_REDIS_PASSWORD")
	conf.Core.Kvothe.Redis.PoolSize = utils.GetInt(os.Getenv("KVOTHE_SERVICE_REDIS_POOLSIZE"))
	conf.Core.Kvothe.Redis.MinIdleConns = utils.GetInt(os.Getenv("KVOTHE_SERVICE_REDIS_MINIDLECONNS"))

	conf.Core.Kvothe.DBPostgres.READ.Username = os.Getenv("KVOTHE_SERVICE_POSTGRES_DATABASE_READ_USERNAME")
	conf.Core.Kvothe.DBPostgres.READ.Password = os.Getenv("KVOTHE_SERVICE_POSTGRES_DATABASE_READ_PASSWORD")
	conf.Core.Kvothe.DBPostgres.READ.URL = os.Getenv("KVOTHE_SERVICE_POSTGRES_DATABASE_READ_URL")
	conf.Core.Kvothe.DBPostgres.READ.Port = os.Getenv("KVOTHE_SERVICE_POSTGRES_DATABASE_READ_PORT")
	conf.Core.Kvothe.DBPostgres.READ.DBName = os.Getenv("KVOTHE_SERVICE_POSTGRES_DATABASE_READ_DB_NAME")
	conf.Core.Kvothe.DBPostgres.READ.MaxIdleConns = utils.GetInt(os.Getenv("KVOTHE_SERVICE_POSTGRES_DATABASE_READ_MAXIDLECONNS"))
	conf.Core.Kvothe.DBPostgres.READ.MaxOpenConns = utils.GetInt(os.Getenv("KVOTHE_SERVICE_POSTGRES_DATABASE_READ_MAXOPENCONNS"))
	conf.Core.Kvothe.DBPostgres.READ.MaxLifeTime = utils.GetInt(os.Getenv("KVOTHE_SERVICE_POSTGRES_DATABASE_READ_MAXLIFETIME"))
	conf.Core.Kvothe.DBPostgres.READ.Timeout = (os.Getenv("KVOTHE_SERVICE_POSTGRES_DATABASE_READ_TIMEOUT"))

	conf.Core.Kvothe.DBPostgres.WRITE.Username = os.Getenv("KVOTHE_SERVICE_POSTGRES_DATABASE_WRITE_USERNAME")
	conf.Core.Kvothe.DBPostgres.WRITE.Password = os.Getenv("KVOTHE_SERVICE_POSTGRES_DATABASE_WRITE_PASSWORD")
	conf.Core.Kvothe.DBPostgres.WRITE.URL = os.Getenv("KVOTHE_SERVICE_POSTGRES_DATABASE_WRITE_URL")
	conf.Core.Kvothe.DBPostgres.WRITE.Port = os.Getenv("KVOTHE_SERVICE_POSTGRES_DATABASE_WRITE_PORT")
	conf.Core.Kvothe.DBPostgres.WRITE.DBName = os.Getenv("KVOTHE_SERVICE_POSTGRES_DATABASE_WRITE_DB_NAME")
	conf.Core.Kvothe.DBPostgres.WRITE.MaxIdleConns = utils.GetInt(os.Getenv("KVOTHE_SERVICE_POSTGRES_DATABASE_WRITE_MAXIDLECONNS"))
	conf.Core.Kvothe.DBPostgres.WRITE.MaxOpenConns = utils.GetInt(os.Getenv("KVOTHE_SERVICE_POSTGRES_DATABASE_WRITE_MAXOPENCONNS"))
	conf.Core.Kvothe.DBPostgres.WRITE.MaxLifeTime = utils.GetInt(os.Getenv("KVOTHE_SERVICE_POSTGRES_DATABASE_WRITE_MAXLIFETIME"))
	conf.Core.Kvothe.DBPostgres.WRITE.Timeout = os.Getenv("KVOTHE_SERVICE_POSTGRES_DATABASE_WRITE_TIMEOUT")

	conf.Core.Kvothe.Name = os.Getenv("KVOTHE_SERVICE")
	conf.Core.Kvothe.Environment = os.Getenv("KVOTHE_SERVICE_ENVIRONMENT")
	conf.Core.Kvothe.Port = os.Getenv("KVOTHE_SERVICE_PORT")

	conf.Core.Kvothe.Slack.WebhookURL = os.Getenv("KVOTHE_SERVICE_WEBHOOK_URL")
	conf.Core.Kvothe.Slack.WebhookChannel = os.Getenv("KVOTHE_SERVICE_WEBHOOK_CHANNEL")
	conf.Core.Kvothe.Slack.IsEnableSlack = utils.GetBool(os.Getenv("KVOTHE_SERVICE_IS_ENABLE_SLACK"))

	conf.Core.Kvothe.GRPC.Port = os.Getenv("KVOTHE_SERVICE_GRPC_PORT")
	return &conf
}
