package utils

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// HandleRequestID ...
func HandleRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		requestID := c.Request.Header.Get("X-Request-Id")

		// Create request id with UUID4
		if requestID == "" {
			uuid4 := uuid.NewV4()
			requestID = uuid4.String()
		}

		// Expose it for use in the application
		c.Set("RequestID", requestID)

		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}
