package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/tigrouland/api/mongo/entities"
)

func Self(c *gin.Context) {
	user := c.MustGet("user").(entities.User)
	c.JSON(200, user)
}
