package middlewares

import (
	"github.com/gin-gonic/gin"
)

const (
	// stsHeader           = "Strict-Transport-Security"
	// stsSubdomainString  = "; includeSubdomains"
	frameOptionsKey         = "X-Frame-Options"
	frameOptionsValue       = "DENY"
	contentTypeOptionsKey   = "X-Content-Type-Options"
	contentTypeOptionsValue = "nosniff"
	xssProtectionKey        = "X-XSS-Protection"
	xssProtectionValue      = "1; mode=block"
	// cspHeader           = "Content-Security-Policy"

	contentTypeDefaultKey = "Content-Type"
	contentTypeJSONValue  = "application/json"
	xPoweredByKey         = "x-powered-by"
	xPoweredByValue       = "golang"
)

func SecureMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header(contentTypeOptionsKey, contentTypeOptionsValue)
		c.Header(frameOptionsKey, frameOptionsValue)
		c.Header(xssProtectionKey, xssProtectionValue)
		c.Header(contentTypeDefaultKey, contentTypeJSONValue)
		c.Header(xPoweredByKey, xPoweredByValue)
		c.Next()
	}
}
