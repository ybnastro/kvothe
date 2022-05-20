package controllers

import (
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/astronautsid/astro-ims-be/infrastructures"
	"github.com/astronautsid/astro-ims-be/interfaces"
	"github.com/astronautsid/astro-ims-be/resources"
	"github.com/astronautsid/astro-ims-be/utils"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	Services       interfaces.InterfaceInventoryService
	Configurations *resources.AppConfig
	Slack          interfaces.ISlackWebhook
	SlackConfig    *infrastructures.SlackWebhook
	PanicHandler   interfaces.IPanicHandler
	HTTP           interfaces.IHTTP
}

func (isc *InventoryController) TriggerPanic(c *gin.Context) {
	panic("panic triggered via endpoint /api/panic")
}

func (isc *InventoryController) PingHandler(c *gin.Context) {
	utils.ResponseJSON(c, "PONG")
	return
}

func (isc *InventoryController) SlackManualHandler(c *gin.Context) {
	name := c.Query("name")
	errs := isc.Slack.Send(
		isc.Configurations.Core.Inventory.Slack.WebhookURL,
		"",
		slack.Payload{
			Text:      fmt.Sprintf("Hello %s", name),
			Username:  "inventory backend",
			Channel:   isc.Configurations.Core.Inventory.Slack.WebhookChannel,
			IconEmoji: ":monkey_face:",
		},
	)
	if len(errs) > 0 {
		utils.RespondWithError(c, http.StatusInternalServerError, errs)
		return
	}
	utils.ResponseJSON(c, "OK")
}
