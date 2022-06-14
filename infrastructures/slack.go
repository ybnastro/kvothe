package infrastructures

import (
	"context"

	"github.com/slack-go/slack"
)

type Slack struct {
	SlackClient *slack.Client
}

func NewSlackClient(token string) *Slack {
	client := slack.New(token)

	s := &Slack{}
	s.SlackClient = client

	return s
}

func (s *Slack) New(token string, option ...slack.Option) *slack.Client {
	s.SlackClient = slack.New(token)

	return s.SlackClient
}

func (s *Slack) UploadFile(params slack.FileUploadParameters) (file *slack.File, err error) {
	return s.SlackClient.UploadFileContext(context.Background(), params)
}

func (s *Slack) SendSlack(channelID string, options ...slack.MsgOption) (string, string, error) {
	return s.SlackClient.PostMessage(channelID, options...)
}
