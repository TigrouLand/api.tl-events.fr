package entities

import (
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	ID          int64            `json:"id" bson:"id"`
	Name        string           `bson:"name" json:"name"`
	Rank        string           `bson:"rank" json:"rank"`
	UUID        primitive.Binary `bson:"uuid" json:"-"`
	DecodedUUID uuid.UUID        `json:"uuid"`
	Kills       int16            `bson:"kills" json:"kills"`
	Deaths      int16            `bson:"deaths" json:"deaths"`
	Wins        int16            `bson:"wins" json:"wins"`
}

func (p *Player) DecodeUUID() error {
	if p.UUID.IsZero() {
		return errors.New("uuid is zero")
	}
	playerUUID, err := uuid.FromBytes(p.UUID.Data)
	if err != nil {
		return err
	}
	p.DecodedUUID = playerUUID
	return nil
}
