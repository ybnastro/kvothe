package interfaces

import (
	"github.com/gin-gonic/gin"

	"github.com/SurgicalSteel/kvothe/resources"
)

// IRouter is interface for routing
type IRouter interface {
	Routing(config *resources.AppConfig, dbObj map[string]IDatabase, redisdb map[string]IRedis) *gin.Engine
}
