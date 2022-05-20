package interfaces

import (
	"github.com/gin-gonic/gin"

	"github.com/astronautsid/astro-ims-be/resources"
)

// IRouter is interface for routing
type IRouter interface {
	Routing(config *resources.AppConfig, dbObj map[string]IDatabase, redisdb map[string]IRedis) *gin.Engine
}
