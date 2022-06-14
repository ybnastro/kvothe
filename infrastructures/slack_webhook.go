package infrastructures

import (
	"context"
	"errors"

	"github.com/ashwanthkumar/slack-go-webhook"
	slackAPI "github.com/slack-go/slack"
)

type SlackWebhookFuncOption func(*SlackWebhook)

func WithWebhook(url string, channel string) SlackWebhookFuncOption {
	return func(sw *SlackWebhook) {
		sw.SlackWebhookURL = url
		sw.SlackWebhookChannel = channel
	}
}

type SlackWebhook struct {
	SlackWebhookEnv     string
	SlackWebhookURL     string
	SlackWebhookChannel string

	SlackToken          string
	SlackUploadFilePath string
}

func NewSlackWebhook(opts ...SlackWebhookFuncOption) *SlackWebhook {
	return new(SlackWebhook).Assign(opts...)
}

func (sw *SlackWebhook) Assign(opts ...SlackWebhookFuncOption) *SlackWebhook {
	for _, opt := range opts {
		opt(sw)
	}
	return sw
}

func (sw *SlackWebhook) validateWebhook(payload *slackAPI.WebhookMessage) (err error) {
	if payload == nil {
		return
	}
	if sw.SlackWebhookURL == "" {
		err = errors.New("the webhook url musn't be empty")
		return
	}
	if sw.SlackWebhookChannel == "" {
		err = errors.New("the channel mustn't be empty")
		return
	}

	payload.Channel = sw.SlackWebhookChannel
	return
}

func (sw *SlackWebhook) SendWebhook(ctx context.Context, payload slackAPI.WebhookMessage) (err error) {
	err = sw.validateWebhook(&payload)
	if err != nil {
		return
	}

	err = slackAPI.PostWebhook(sw.SlackWebhookURL, &payload)
	return
}

// sending report image and text
// Send ...
func (s *SlackWebhook) Send(webhookURL, proxy string, payload slack.Payload) []error {

	if s == nil {
		return []error{
			errors.New("invalid initialize data"),
		}
	}

	errData := make([]error, 0)
	var err error
	if webhookURL == "" {
		err = errors.New("[SendWebhookSlack] webhook url should not empty")
		errData = append(errData, err)
		return errData
	}

	if payload.Channel == "" {
		err = errors.New("[SendWebhookSlack] channel url should not empty")
		errData = append(errData, err)
		return errData
	}

	errSlack := slack.Send(webhookURL, proxy, payload)
	if len(errSlack) > 0 {
		return errSlack
	}
	return nil
}
