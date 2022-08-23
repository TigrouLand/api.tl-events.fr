package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tigrouland/api/mongo"
	"github.com/tigrouland/api/mongo/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func Members(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := mongo.Get().Collection("players").Find(ctx, bson.D{}, options.Find().SetSort(bson.D{{"wins", -1}, {"kills", -1}}))
	if err != nil {
		log.Fatal("an error occurred while retrieving the players' data: ", err)
	}
	var players []entities.Player
	for cursor.Next(ctx) {
		var player entities.Player
		err = cursor.Decode(&player)
		if err != nil {
			log.Fatal(err)
		}
		var playerUUID uuid.UUID
		if !player.UUID.IsZero() {
			playerUUID, err = uuid.FromBytes(player.UUID.Data)
			if err != nil {
				log.Fatal(err)
			}
			player.DecodedUUID = playerUUID
		}
		players = append(players, player)
	}
	c.JSON(200, players)
}
