package server

import (
	"github.com/astronautsid/astro-ims-be/interfaces"
	"github.com/astronautsid/astro-ims-be/middlewares"
	"github.com/astronautsid/astro-ims-be/resources"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

//route is a struct for router
type route struct {
	interfaces.IRouter
}

var r *route

// Router is a func to initialize route struct
func Router() interfaces.IRouter {
	if r == nil {
		r = &route{}
	}
	return r
}

//Routing is a function for http routing
func (route *route) Routing(config *resources.AppConfig, dbObj map[string]interfaces.IDatabase, redisdb map[string]interfaces.IRedis) *gin.Engine {

	inventoryController := ServiceInject(config, dbObj, redisdb)

	gin.SetMode(config.GINMode)
	var engine *gin.Engine
	if config.GINMode == resources.Release {
		engine = gin.New()
	} else {
		engine = gin.Default()
	}

	engine.Use(gzip.Gzip(gzip.BestCompression))
	engine.Use(middlewares.SecureMiddleware())
	engine.Use(middlewares.PanicGlobalRecover("inventory-service", inventoryController))

	noAuth := engine.Group("/api")
	{
		noAuth.GET("/ping", inventoryController.PingHandler)
		noAuth.GET("/panic", inventoryController.TriggerPanic)
		noAuth.GET("/slack", inventoryController.SlackManualHandler)
	}

	return engine
}
