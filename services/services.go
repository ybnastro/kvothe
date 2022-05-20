//package services defines the app dependency and route mapping
package services

import (
	"github.com/astronautsid/astro-ims-be/interfaces"
	"github.com/astronautsid/astro-ims-be/resources"
)

type InventoryService struct {
	Repo    interfaces.InterfaceInventoryRepository
	HTTP    interfaces.IHTTP
	Conf    *resources.AppConfig
	Context interfaces.IContext
}
