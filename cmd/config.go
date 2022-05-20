package main

import (
	"log"
	"os"

	"github.com/astronautsid/astro-ims-be/resources"
	"github.com/astronautsid/astro-ims-be/utils"

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

	conf.Core.Inventory.Redis.URL = os.Getenv("INVENTORY_SERVICE_REDIS_URL")
	conf.Core.Inventory.Redis.Port = utils.GetInt(os.Getenv("INVENTORY_SERVICE_REDIS_PORT"))
	conf.Core.Inventory.Redis.DB = utils.GetInt(os.Getenv("INVENTORY_SERVICE_REDIS_DB"))
	conf.Core.Inventory.Redis.Password = os.Getenv("INVENTORY_SERVICE_REDIS_PASSWORD")
	conf.Core.Inventory.Redis.PoolSize = utils.GetInt(os.Getenv("INVENTORY_SERVICE_REDIS_POOLSIZE"))
	conf.Core.Inventory.Redis.MinIdleConns = utils.GetInt(os.Getenv("INVENTORY_SERVICE_REDIS_MINIDLECONNS"))

	conf.Core.Inventory.DBPostgres.READ.Username = os.Getenv("INVENTORY_SERVICE_POSTGRES_DATABASE_READ_USERNAME")
	conf.Core.Inventory.DBPostgres.READ.Password = os.Getenv("INVENTORY_SERVICE_POSTGRES_DATABASE_READ_PASSWORD")
	conf.Core.Inventory.DBPostgres.READ.URL = os.Getenv("INVENTORY_SERVICE_POSTGRES_DATABASE_READ_URL")
	conf.Core.Inventory.DBPostgres.READ.Port = os.Getenv("INVENTORY_SERVICE_POSTGRES_DATABASE_READ_PORT")
	conf.Core.Inventory.DBPostgres.READ.DBName = os.Getenv("INVENTORY_SERVICE_POSTGRES_DATABASE_READ_DB_NAME")
	conf.Core.Inventory.DBPostgres.READ.MaxIdleConns = utils.GetInt(os.Getenv("INVENTORY_SERVICE_POSTGRES_DATABASE_READ_MAXIDLECONNS"))
	conf.Core.Inventory.DBPostgres.READ.MaxOpenConns = utils.GetInt(os.Getenv("INVENTORY_SERVICE_POSTGRES_DATABASE_READ_MAXOPENCONNS"))
	conf.Core.Inventory.DBPostgres.READ.MaxLifeTime = utils.GetInt(os.Getenv("INVENTORY_SERVICE_POSTGRES_DATABASE_READ_MAXLIFETIME"))
	conf.Core.Inventory.DBPostgres.READ.Timeout = (os.Getenv("INVENTORY_SERVICE_POSTGRES_DATABASE_READ_TIMEOUT"))

	conf.Core.Inventory.DBPostgres.WRITE.Username = os.Getenv("INVENTORY_SERVICE_POSTGRES_DATABASE_WRITE_USERNAME")
	conf.Core.Inventory.DBPostgres.WRITE.Password = os.Getenv("INVENTORY_SERVICE_POSTGRES_DATABASE_WRITE_PASSWORD")
	conf.Core.Inventory.DBPostgres.WRITE.URL = os.Getenv("INVENTORY_SERVICE_POSTGRES_DATABASE_WRITE_URL")
	conf.Core.Inventory.DBPostgres.WRITE.Port = os.Getenv("INVENTORY_SERVICE_POSTGRES_DATABASE_WRITE_PORT")
	conf.Core.Inventory.DBPostgres.WRITE.DBName = os.Getenv("INVENTORY_SERVICE_POSTGRES_DATABASE_WRITE_DB_NAME")
	conf.Core.Inventory.DBPostgres.WRITE.MaxIdleConns = utils.GetInt(os.Getenv("INVENTORY_SERVICE_POSTGRES_DATABASE_WRITE_MAXIDLECONNS"))
	conf.Core.Inventory.DBPostgres.WRITE.MaxOpenConns = utils.GetInt(os.Getenv("INVENTORY_SERVICE_POSTGRES_DATABASE_WRITE_MAXOPENCONNS"))
	conf.Core.Inventory.DBPostgres.WRITE.MaxLifeTime = utils.GetInt(os.Getenv("INVENTORY_SERVICE_POSTGRES_DATABASE_WRITE_MAXLIFETIME"))
	conf.Core.Inventory.DBPostgres.WRITE.Timeout = os.Getenv("INVENTORY_SERVICE_POSTGRES_DATABASE_WRITE_TIMEOUT")

	conf.Core.Inventory.Name = os.Getenv("INVENTORY_SERVICE")
	conf.Core.Inventory.Environment = os.Getenv("INVENTORY_SERVICE_ENVIRONMENT")
	conf.Core.Inventory.Port = os.Getenv("INVENTORY_SERVICE_PORT")

	conf.Core.Inventory.Slack.WebhookURL = os.Getenv("INVENTORY_SERVICE_WEBHOOK_URL")
	conf.Core.Inventory.Slack.WebhookChannel = os.Getenv("INVENTORY_SERVICE_WEBHOOK_CHANNEL")
	conf.Core.Inventory.Slack.IsEnableSlack = utils.GetBool(os.Getenv("INVENTORY_SERVICE_IS_ENABLE_SLACK"))
	log.Println(conf.Core.Inventory.Slack)
	return &conf
}
