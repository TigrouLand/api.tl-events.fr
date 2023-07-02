package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/tigrouland/api/mongo/entities"
)

func UserRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawUser, exists := c.Get("user")
		if !exists {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		user := rawUser.(entities.User)
		if (user == entities.User{}) {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
