package entities

type Team struct {
	Name   string `bson:"name" json:"name"`
	Colors []int  `bson:"colors" json:"colors"`
}
