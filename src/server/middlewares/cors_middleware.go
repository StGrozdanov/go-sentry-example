package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CORS handles the CORS headers and origins
//
//nolint:lll
func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, connection, Cache-Control, X-Requested-With, X-Authorization, Baggage, Sentry-Trace")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, HEAD, PATCH")

		if ctx.Request.Method == "OPTIONS" {
			ctx.JSON(http.StatusOK, map[string]interface{}{})
			return
		}

		ctx.Next()
	}
}
