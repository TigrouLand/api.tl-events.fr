package auth

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tigrouland/api/core"
	"github.com/tigrouland/api/mongo"
	"github.com/tigrouland/api/mongo/entities"
	"go.mongodb.org/mongo-driver/bson"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"os"
	"strconv"
)

type DiscordRessourceOwner struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func Callback(c *gin.Context) {
	session := sessions.Default(c)
	state := session.Get("state")
	if state != c.Query("state") {
		c.JSON(400, gin.H{"error": "invalid state"})
		return
	}

	token, err := core.DiscordOAuth.Exchange(context.Background(), c.Query("code"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid code"})
		return
	}

	res, err := core.DiscordOAuth.Client(context.Background(), token).Get("https://discord.com/api/users/@me")
	if err != nil || res.StatusCode != 200 {
		c.JSON(500, gin.H{"error": "an error occurred while retrieving user information"})
		return
	}

	defer res.Body.Close()

	var discordUser DiscordRessourceOwner
	err = json.NewDecoder(res.Body).Decode(&discordUser)
	if err != nil {
		c.JSON(500, gin.H{"error": "an error occurred while decoding user information"})
		return
	}

	discordID, err := strconv.Atoi(discordUser.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "an error occurred while converting user id"})
		return
	}

	var player entities.Player
	err = mongo.Get().Collection("players").FindOne(context.Background(), bson.M{"id": discordID}).Decode(&player)
	if err != nil {
		if errors.Is(err, mongo2.ErrNoDocuments) {
			c.Redirect(302, os.Getenv("FRONTEND_URL")+"?error=not_registered")
			return
		}
		c.JSON(500, gin.H{"error": "an error occurred while retrieving user information"})
		return
	}

	_, err = mongo.Get().Collection("players").UpdateOne(
		context.Background(),
		bson.M{"id": discordID},
		bson.M{"$set": bson.A{
			bson.M{"oauth.token": token.AccessToken},
			bson.M{"oauth.refreshToken": token.RefreshToken},
			bson.M{"oauth.expiry": token.Expiry},
		}},
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "an error occurred while updating user information"})
		return
	}

	session.Set("discordID", discordUser.ID)
	err = session.Save()
	if err != nil {
		c.JSON(500, gin.H{"error": "an error occurred while saving the session"})
		return
	}

	c.Redirect(302, os.Getenv("FRONTEND_URL"))
}
