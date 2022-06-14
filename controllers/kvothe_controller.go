package controllers

import (
	"errors"
	"strconv"

	"github.com/SurgicalSteel/kvothe/infrastructures"
	"github.com/SurgicalSteel/kvothe/interfaces"
	"github.com/SurgicalSteel/kvothe/resources"
	"github.com/SurgicalSteel/kvothe/utils"
	"github.com/ashwanthkumar/slack-go-webhook"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KvotheController struct {
	Services       interfaces.InterfaceKvotheService
	Configurations *resources.AppConfig
	Slack          interfaces.ISlackWebhook
	SlackConfig    *infrastructures.SlackWebhook
	PanicHandler   interfaces.IPanicHandler
	HTTP           interfaces.IHTTP
}

func (kc *KvotheController) TriggerPanic(c *gin.Context) {
	panic("panic triggered via endpoint /api/panic")
}

func (kc *KvotheController) PingHandler(c *gin.Context) {
	utils.ResponseJSON(c, "PONG")
	return
}

func (kc *KvotheController) SlackManualHandler(c *gin.Context) {
	name := c.Query("name")
	errs := kc.Slack.Send(
		kc.Configurations.Core.Kvothe.Slack.WebhookURL,
		"",
		slack.Payload{
			Text:      fmt.Sprintf("Hello %s", name),
			Username:  "kvothe backend",
			Channel:   kc.Configurations.Core.Kvothe.Slack.WebhookChannel,
			IconEmoji: ":monkey_face:",
		},
	)
	if len(errs) > 0 {
		utils.RespondWithError(c, http.StatusInternalServerError, errs)
		return
	}
	utils.ResponseJSON(c, "OK")
}

func (kc *KvotheController) GetSongQuoteByIDHandler(c *gin.Context) {
	rawID := c.Param("id")
	if len(rawID) == 0 {
		utils.RespondWithError(c, http.StatusBadRequest, errors.New("Song Quote ID is not defined"))
		return
	}

	id, _ := strconv.ParseInt(rawID, 10, 64)
	songQuote, httpStatus, err := kc.Services.GetSongQuoteByID(id)
	if err != nil {
		utils.RespondWithError(c, httpStatus, err.Error())
		return
	}

	utils.ResponseJSON(c, songQuote)
}

func (kc *KvotheController) BackfillRedisHandler(c *gin.Context) {
	err := kc.Services.BackfillRedis()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseJSON(c, "OK")
}

func (kc *KvotheController) GetAllSongPage(c *gin.Context) {
	data, httpStatus, err := kc.Services.GetAllSongData()
	if err != nil {
		utils.RespondWithError(c, httpStatus, err.Error())
		return
	}

	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	c.HTML(http.StatusOK, "all_song_quote.tmpl", gin.H{
		"title":     "All Song Quote",
		"quoteData": data,
	})

}
