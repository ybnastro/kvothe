package infrastructures

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SurgicalSteel/kvothe/interfaces"
	"github.com/SurgicalSteel/kvothe/resources"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/gin-gonic/gin"
)

type PanicHandlerController struct {
	Slack       interfaces.ISlackWebhook
	SlackConfig *SlackWebhook
	Logger      interfaces.ILogger
}

func (pc *PanicHandlerController) SetPanicResp(c *gin.Context, serviceName, message string) *resources.PanicHandlerResponse {

	resp := new(resources.PanicHandlerResponse)

	resp.Message = message
	resp.Service = serviceName
	resp.Status = "ERROR"
	resp.StatusCode = http.StatusInternalServerError

	if c != nil {
		resp.Method = c.Request.Method
		context := fmt.Sprintf("%+v", c.Request)
		resp.Context = context
		url := fmt.Sprintf("%s%s", c.Request.Host, c.Request.URL)
		resp.URL = url
	}

	resp.Env = pc.SlackConfig.SlackWebhookEnv

	return resp
}

func (pc *PanicHandlerController) SetPayloadSlack(resp *resources.PanicHandlerResponse) slack.Payload {

	color := "FF0000"
	attachment := slack.Attachment{
		Color: &color,
	}
	attachment1 := slack.Attachment{
		Color: &color,
	}

	// fields request context
	attachment.AddField(slack.Field{Title: "URL", Value: resp.URL})

	// fields panic
	startedAt := time.Now().Format("2006-01-02 15:04:05")
	attachment1.AddField(slack.Field{Title: "Panic-service", Value: resp.Service}).AddField(slack.Field{Title: "Env", Value: pc.SlackConfig.SlackWebhookEnv})
	attachment1.AddField(slack.Field{Title: "Started At", Value: startedAt}).AddField(slack.Field{Title: "Created by", Value: "kvothe backend"})

	payload := slack.Payload{
		Text:        resp.Message,
		Username:    "kvothe backend",
		Channel:     pc.SlackConfig.SlackWebhookChannel,
		IconEmoji:   ":monkey_face:",
		Attachments: []slack.Attachment{attachment, attachment1},
	}

	return payload

}

func (pc *PanicHandlerController) GetPanicAndSendToSlack(c *gin.Context, serviceName, message string) error {

	err := errors.New("service unavailable")
	if message == "" {
		return err
	}

	if pc.Slack == nil || pc.SlackConfig == nil {
		return err
	}

	resp := new(resources.PanicHandlerResponse)
	resp = pc.SetPanicResp(c, serviceName, message)

	payloaddata := pc.SetPayloadSlack(resp)

	webhookURL := pc.SlackConfig.SlackWebhookURL
	errSlack := pc.Slack.Send(webhookURL, "", payloaddata)
	if len(errSlack) > 0 {
		log.Println(errSlack[0])
		return errSlack[0]
	}
	return nil
}
