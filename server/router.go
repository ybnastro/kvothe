package server

import (
	"github.com/SurgicalSteel/kvothe/interfaces"
	"github.com/SurgicalSteel/kvothe/middlewares"
	"github.com/SurgicalSteel/kvothe/resources"

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

	kvotheController := ServiceInject(config, dbObj, redisdb)

	gin.SetMode(config.GINMode)
	var engine *gin.Engine
	if config.GINMode == resources.Release {
		engine = gin.New()
	} else {
		engine = gin.Default()
	}

	engine.Use(gzip.Gzip(gzip.BestCompression))
	engine.Use(middlewares.SecureMiddleware())
	engine.Use(middlewares.PanicGlobalRecover("kvothe-service", kvotheController))

	engine.LoadHTMLGlob("../files/templates/*.tmpl")
	noAuth := engine.Group("/api")
	{
		noAuth.GET("/ping", kvotheController.PingHandler)
		noAuth.GET("/panic", kvotheController.TriggerPanic)
		noAuth.GET("/slack", kvotheController.SlackManualHandler)
		noAuth.GET("/quote/:id", kvotheController.GetSongQuoteByIDHandler)
		noAuth.POST("/backfill-redis", kvotheController.BackfillRedisHandler)
	}

	noAuthPage := engine.Group("/page")
	{
		noAuthPage.GET("/all", kvotheController.GetAllSongPage)
	}

	return engine
}
