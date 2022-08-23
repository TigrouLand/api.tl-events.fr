package entities

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	Name        string           `bson:"name" json:"name"`
	Rank        string           `bson:"rank" json:"rank"`
	UUID        primitive.Binary `bson:"uuid" json:"-"`
	DecodedUUID uuid.UUID        `json:"uuid"`
	Kills       int16            `bson:"kills" json:"kills"`
	Deaths      int16            `bson:"deaths" json:"deaths"`
	Wins        int16            `bson:"wins" json:"wins"`
}
