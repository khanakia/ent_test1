// Package jutil provide some common sets of shared function
package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Throw Malfunction error for Gin Request
func GinErrorMalfunction(c *gin.Context) {
	err := &ResponseError{
		Message: "Malfunctioned request.",
	}

	err.Send(c, http.StatusUnauthorized)
	// c.JSON(http.StatusBadRequest, gin.H{"error": error})
}

// Throw Bad Request error for Gin Request
func GinErrorBadRequest(c *gin.Context) {
	err := &ResponseError{
		Message: "Bad request.",
	}

	err.Send(c, http.StatusUnauthorized)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-captcha-res, X-CAPTCHA-BYPASS, x-auth-token, token, Website")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
