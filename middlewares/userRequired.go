package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/tigrouland/api/core"
)

func UserRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawUser, exists := c.Get("user")
		if !exists {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		user := rawUser.(core.User)
		if (user == core.User{}) {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
