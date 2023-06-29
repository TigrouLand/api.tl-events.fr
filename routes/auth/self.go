package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/tigrouland/api/core"
)

func Self(c *gin.Context) {
	user := c.MustGet("user").(core.User)
	c.JSON(200, user)
}
