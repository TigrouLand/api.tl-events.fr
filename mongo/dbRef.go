package mongo

type DBRef struct {
	Ref interface{} `bson:"$ref"`
	ID  interface{} `bson:"$id"`
	DB  interface{} `bson:"$db"`
}
