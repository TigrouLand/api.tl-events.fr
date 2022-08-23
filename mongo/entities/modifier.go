package entities

type Modifier struct {
	Name        string   `bson:"name" json:"name"`
	Enabled     bool     `bson:"enabled" json:"enabled"`
	Material    string   `json:"material" bson:"material"`
	Description []string `json:"description" bson:"description"`
	Web         []string `json:"web,omitempty" bson:"web"`
}
