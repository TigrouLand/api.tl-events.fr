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

func GetGames() []entities.Game {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := mongo.Get().Collection("games").Find(ctx, bson.M{"archiveDate": -1}, options.Find().SetSort(bson.D{{"startDate", -1}}))
	if err != nil {
		log.Fatal("an error occurred while retrieving the games: ", err)
	}
	games := []entities.Game{}
	for cursor.Next(ctx) {
		var game entities.Game
		err = cursor.Decode(&game)
		if err != nil {
			log.Fatal(err)
		}

		DecodeGame(&game, ctx)
		games = append(games, game)
	}
	return games
}

func GetArchivedGames() []entities.ArchivedGame {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := mongo.Get().Collection("games").Find(ctx, bson.M{"archiveDate": bson.M{"$ne": -1}}, options.Find().SetSort(bson.D{{"startDate", -1}}))
	if err != nil {
		log.Fatal("an error occurred while retrieving the games: ", err)
	}
	games := []entities.ArchivedGame{}
	for cursor.Next(ctx) {
		var game entities.ArchivedGame
		err = cursor.Decode(&game)
		if err != nil {
			log.Fatal(err)
		}

		DecodeArchivedGame(&game, ctx)
		games = append(games, game)
	}
	return games
}

func Games(c *gin.Context) {
	c.JSON(200, bson.M{"games": GetGames(), "archivedGames": GetArchivedGames()})
}

func DecodeGame(game *entities.Game, ctx context.Context) {
	if len(game.Players) > 0 {
		for i := range game.Players {
			var player entities.Player
			var playerUUID uuid.UUID
			if !(game.Players[i] == mongo.DBRef{}) {
				result := mongo.Get().Collection("players").FindOne(ctx, bson.M{"_id": game.Players[i].ID})
				if result.Err() != nil {
					log.Fatal(result.Err())
				}
				err := result.Decode(&player)
				if err != nil {
					log.Fatal(err)
				}

				if !player.UUID.IsZero() {
					playerUUID, err = uuid.FromBytes(player.UUID.Data)
					if err != nil {
						log.Fatal(err)
					}

					game.PlayersUUID = append(game.PlayersUUID, playerUUID)
				}
			}
		}

		//if len(game.TeamRefs) > 0 {
		//	for i := range game.TeamRefs {
		//		var team entities.Team
		//		if !(game.TeamRefs[i] == mongo.DBRef{}) {
		//			result := mongo.Get().Collection("teams").FindOne(ctx, bson.M{"_id": game.TeamRefs[i].ID})
		//			if result.Err() != nil {
		//				log.Fatal(result.Err())
		//			}
		//			err := result.Decode(&team)
		//			if err != nil {
		//				log.Fatal(err)
		//			}
		//
		//			game.Teams = append(game.Teams, team)
		//		}
		//	}
		//}
	}
}

func DecodeArchivedGame(game *entities.ArchivedGame, ctx context.Context) {
	if len(game.Players) > 0 {
		for i := range game.Players {
			var player entities.Player
			var playerUUID uuid.UUID
			if !(game.Players[i] == mongo.DBRef{}) {
				result := mongo.Get().Collection("players").FindOne(ctx, bson.M{"_id": game.Players[i].ID})
				if result.Err() != nil {
					log.Fatal(result.Err())
				}
				err := result.Decode(&player)
				if err != nil {
					log.Fatal(err)
				}

				if !player.UUID.IsZero() {
					playerUUID, err = uuid.FromBytes(player.UUID.Data)
					if err != nil {
						log.Fatal(err)
					}

					game.PlayersUUID = append(game.PlayersUUID, playerUUID)
				}
			}
		}
	}

	if len(game.Alive) > 0 {
		for i := range game.Alive {
			var player entities.Player
			var playerUUID uuid.UUID
			if !(game.Alive[i] == mongo.DBRef{}) {
				result := mongo.Get().Collection("players").FindOne(ctx, bson.M{"_id": game.Alive[i].ID})
				if result.Err() != nil {
					log.Fatal(result.Err())
				}
				err := result.Decode(&player)
				if err != nil {
					log.Fatal(err)
				}

				if !player.UUID.IsZero() {
					playerUUID, err = uuid.FromBytes(player.UUID.Data)
					if err != nil {
						log.Fatal(err)
					}

					game.AliveUUID = append(game.AliveUUID, playerUUID)
				}
			}
		}
	}

	if len(game.Moles) > 0 {
		for i := range game.Moles {
			var player entities.Player
			var playerUUID uuid.UUID
			if !(game.Moles[i] == mongo.DBRef{}) {
				result := mongo.Get().Collection("players").FindOne(ctx, bson.M{"_id": game.Moles[i].ID})
				if result.Err() != nil {
					log.Fatal(result.Err())
				}
				err := result.Decode(&player)
				if err != nil {
					log.Fatal(err)
				}

				if !player.UUID.IsZero() {
					playerUUID, err = uuid.FromBytes(player.UUID.Data)
					if err != nil {
						log.Fatal(err)
					}

					game.MolesUUID = append(game.MolesUUID, playerUUID)
				}
			}
		}
	}

	if len(game.TeamRefs) > 0 {
		for i := range game.TeamRefs {
			if !(game.TeamRefs[i] == mongo.DBRef{}) {
				var team entities.Team
				result := mongo.Get().Collection("teams").FindOne(ctx, bson.M{"_id": game.TeamRefs[i].ID})
				if result.Err() != nil {
					log.Fatal(result.Err())
				}

				err := result.Decode(&team)
				if err != nil {
					log.Fatal(err)
				}

				game.Teams = append(game.Teams, team)
			}
		}
	}
}
