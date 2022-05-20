package repositories

import (
	"github.com/astronautsid/astro-ims-be/interfaces"
	"github.com/astronautsid/astro-ims-be/resources"
)

type InventoryRepository struct {
	DB      interfaces.IDatabase
	Redis   interfaces.IRedis
	HTTP    interfaces.IHTTP
	Conf    *resources.AppConfig
	Context interfaces.IContext
}
