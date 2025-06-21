package middlewares

import (
	"github.com/gin-gonic/gin"
)

func CheckPerm(resource string, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
