package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tigrouland/api/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

func Stats(c *gin.Context) {
	database := mongo.Get()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	membersCount, err := database.Collection("players").CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	modifiersCount, err := database.Collection("modifiers").CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	gamesCount, err := database.Collection("games").CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, bson.M{
		"members":   membersCount,
		"modifiers": modifiersCount,
		"games":     gamesCount,
	})
}
