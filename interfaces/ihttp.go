package interfaces

//go:generate mockgen --destination=../mocks/mock_ihttp.go --package=mocks --source=ihttp.go
import "github.com/SurgicalSteel/kvothe/resources"

//IHTTP is the general interface for HTTP Call
type IHTTP interface {
	CallService(method, url string, requestBody []byte) (string, *resources.ApplicationError)
	CallServiceByte(method, url string, requestBody []byte) ([]byte, *resources.ApplicationError)
}
