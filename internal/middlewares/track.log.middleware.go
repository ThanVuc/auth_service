package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/*
	@Author: Sinh
	@Date:
	@Description: Middleware to track logs by generating a unique request ID if not provided by
	the proxy
*/

func TrackLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.GetHeader("x-request-id")
		if requestId == "" {
			requestId = uuid.New().String()
		}
		c.Set("requestId", requestId)
		c.Next()
	}
}
