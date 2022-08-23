package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var mongoDatabase *mongo.Database

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("an error occurred while connecting to mongodb: ", err)
	}
	mongoDatabase = client.Database(os.Getenv("MONGO_DATABASE"))
}

func Get() *mongo.Database {
	return mongoDatabase
}
