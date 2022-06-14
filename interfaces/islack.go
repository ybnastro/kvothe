package interfaces

//go:generate mockgen --destination=../mocks/mock_islack.go --package=mocks --source=islack.go
import "github.com/slack-go/slack"

// ISlack used for slack api interaction
type ISlack interface {
	New(token string, option ...slack.Option) *slack.Client
	UploadFile(params slack.FileUploadParameters) (file *slack.File, err error)
	// SendSlack ...
	SendSlack(channelID string, options ...slack.MsgOption) (string, string, error)
}
