package interfaces

//go:generate mockgen --destination=../mocks/mock_ipanic_handler.go --package=mocks --source=ipanic_handler.go
import (
	"github.com/SurgicalSteel/kvothe/resources"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/gin-gonic/gin"
)

//IPanicHandler is a interface for panic handler
type IPanicHandler interface {
	SetPanicResp(c *gin.Context, serviceName, message string) *resources.PanicHandlerResponse
	SetPayloadSlack(resp *resources.PanicHandlerResponse) slack.Payload
	GetPanicAndSendToSlack(c *gin.Context, serviceName, message string) error
}
