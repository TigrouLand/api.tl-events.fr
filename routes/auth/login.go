package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tigrouland/api/core"
	"golang.org/x/oauth2"
)

func Login(c *gin.Context) {
	state, err := core.GenerateRandomString(32)
	if err != nil {
		c.JSON(500, gin.H{"error": "an error occurred while generating the state string"})
		return
	}

	session := sessions.Default(c)
	session.Set("state", state)
	err = session.Save()
	if err != nil {
		c.JSON(500, gin.H{"error": "an error occurred while saving the session"})
		return
	}

	loginUrl := core.DiscordOAuth.AuthCodeURL(state, oauth2.ApprovalForce)
	c.JSON(200, gin.H{
		"loginUrl": loginUrl,
	})
}
