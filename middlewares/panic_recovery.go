package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"
	"sync"

	"github.com/astronautsid/astro-ims-be/controllers"
	"github.com/astronautsid/astro-ims-be/interfaces"
	"github.com/astronautsid/astro-ims-be/utils"

	"github.com/gin-gonic/gin"
)

const (
	inventoryService = "inventory-service"
)

func SendPanicSlackWebhook(ipanic interfaces.IPanicHandler, c *gin.Context, serviceName, message string, done chan bool) {
	err := ipanic.GetPanicAndSendToSlack(c, serviceName, message)
	if err != nil {
		log.Printf("[ERROR] GetPanicAndSendToSlack %s\n", err.Error())
		done <- false
	}
	done <- true
}

func definePanicHandler(serviceName string, r interface{}) interfaces.IPanicHandler {
	switch serviceName {
	case inventoryService:
		return r.(*controllers.InventoryController).PanicHandler
	default:
		return nil
	}

}

func PanicGlobalRecover(serviceName string, r interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func(c *gin.Context) {
			if rec := recover(); rec != nil {

				// that recovery also handle XHR's
				// you need handle it
				message := string(debug.Stack())
				ipanic := definePanicHandler(serviceName, r)

				if ipanic != nil {

					totalRunningPanicRecover := 1

					var wg sync.WaitGroup
					wg.Add(totalRunningPanicRecover)

					doneSlackWebhook := make(chan bool)

					var isSendSlackWebhook bool

					go SendPanicSlackWebhook(ipanic, c, serviceName, message, doneSlackWebhook)
					wg.Done()

					wg.Wait()

					isSendSlackWebhook = <-doneSlackWebhook

					close(doneSlackWebhook)

					if isSendSlackWebhook {
						utils.RespondWithError(c, http.StatusInternalServerError, "service unavailable")
					} else {
						utils.RespondWithError(c, http.StatusBadRequest, "service unavailable bad request")
					}
				}
			}
		}(c)
		c.Next()
	}
}
