package middlewares

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tigrouland/api/mongo"
	"github.com/tigrouland/api/mongo/entities"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func User() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("discordId")
		var player entities.Player
		switch userID.(type) {
		case uint:
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			err := mongo.Get().Collection("players").FindOne(ctx, bson.M{"id": userID.(uint)}).Decode(&player)
			if err == nil {
				player.DecodeUUID()
			}
		}
		c.Next()
	}
}
