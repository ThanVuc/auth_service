package middlewares

import (
	"auth_service/global"
	"auth_service/pkg/response"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

/*
@Author: Sinh
@Date:
@Description: Global error handler middleware for the application.
*/
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger := global.Logger
				requestIdVal, _ := c.Get("requestId")
				requestId, _ := requestIdVal.(string)
				switch e := r.(type) {
				case response.ErrorResponse:
					// Needing more details ? -> add debug.Stack() to the logger
					if e.StatusCode >= 500 {
						logger.Error(e, requestId, debug.Stack())
					} else {
						logger.Error(e, requestId, nil)
					}

					c.JSON(e.StatusCode, e)
				default:
					logger.Error(response.InternalServerError("Unknown panic"), requestId, debug.Stack())
					c.JSON(500, response.AnotherError(http.StatusInternalServerError, "Unknown panic"))
				}
			}

			c.Abort()
		}()
		c.Next()
	}
}
