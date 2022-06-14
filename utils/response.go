package utils

import (
	"fmt"
	"net/http"

	"github.com/SurgicalSteel/kvothe/resources"

	"github.com/gin-gonic/gin"
)

//GetErrorResponse is a function for build errorMessage
func GetErrorResponse(message string, statusCode int, code string) *resources.ApplicationError {
	return &resources.ApplicationError{
		Message:    fmt.Sprintf(message),
		StatusCode: statusCode,
		Code:       code,
	}
}

//RespondWithError for sending response with status not 200
func RespondWithError(c *gin.Context, status int, data interface{}) {
	c.Set("responseBody", data)
	c.JSON(status, data)
}

// ResponseJSON for sending response with status 200
func ResponseJSON(c *gin.Context, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.Set("responseBody", data)
	c.JSON(http.StatusOK, data)
}
