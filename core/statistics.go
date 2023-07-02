package core

import (
	"github.com/tigrouland/api/mongo"
	"github.com/tigrouland/api/mongo/entities"
	"go.mongodb.org/mongo-driver/bson"
)

type UserStats struct {
	Kills  int16 `json:"kills"`
	Deaths int16 `json:"deaths"`
	Wins   int16 `json:"wins"`
}

func GetStatistics(username string) (*UserStats, error) {
	var player entities.Player
	err := mongo.Get().Collection("players").FindOne(nil, bson.M{"name": username}).Decode(&player)
	if err != nil {
		return nil, err
	}

	return &UserStats{
		Kills:  player.Kills,
		Deaths: player.Deaths,
		Wins:   player.Wins,
	}, nil
}
