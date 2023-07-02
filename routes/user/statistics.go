package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tigrouland/api/core"
	"go.mongodb.org/mongo-driver/mongo"
)

func Statistics(c *gin.Context) {
	username := c.Param("username")

	stats, err := core.GetStatistics(username)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(404, gin.H{"error": "user not found"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, stats)
}
