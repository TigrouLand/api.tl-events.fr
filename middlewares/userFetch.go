package middlewares

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tigrouland/api/mongo"
	"github.com/tigrouland/api/mongo/entities"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

func UserFetch() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		rawUserID := session.Get("discordId")
		if rawUserID == nil {
			c.Next()
			return
		}
		userID := rawUserID.(int)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user entities.User
		err := mongo.Get().Collection("users").FindOne(ctx, bson.M{"id": userID}).Decode(&user)
		if err != nil {
			log.Println(err)
			c.Next()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
