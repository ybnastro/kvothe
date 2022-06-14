package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUtils_RespondWithError(t *testing.T) {
	h := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(h)

	testCases := []struct {
		name                string
		inputGinContext     *gin.Context
		inputHTTPStatusCode int
		inputData           interface{}
	}{
		{
			name:                "error 404",
			inputGinContext:     c,
			inputHTTPStatusCode: http.StatusNotFound,
			inputData:           "Not Found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			RespondWithError(tc.inputGinContext, tc.inputHTTPStatusCode, tc.inputData)

			responseBody, exist := tc.inputGinContext.Get("responseBody")
			if !exist {
				t.Error("responseBody key is not exist")
			} else {
				assert.Equal(t, responseBody, tc.inputData)
			}

		})
	}
}
