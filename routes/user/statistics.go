package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tigrouland/api/core"
	"go.mongodb.org/mongo-driver/mongo"
)

func Statistics(c *gin.Context) {
	id := c.Param("id")
	userUuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid uuid"})
		return
	}

	stats, err := core.GetStatistics(userUuid)
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
