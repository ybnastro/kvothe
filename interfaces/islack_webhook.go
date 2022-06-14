package interfaces

//go:generate mockgen --destination=../mocks/mock_islack_webhook.go --package=mocks --source=islack_webhook.go

import (
	"context"

	slack "github.com/ashwanthkumar/slack-go-webhook"
	slackAPI "github.com/slack-go/slack"
)

// ISlackWebhook for sending slack webhook interface
type ISlackWebhook interface {
	SendWebhook(ctx context.Context, payload slackAPI.WebhookMessage) (err error)
	Send(webhookURL, proxy string, payload slack.Payload) []error
}
