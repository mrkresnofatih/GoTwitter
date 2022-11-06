package applications

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var mongoInstance *mongo.Database

func GetMongoInstance() *mongo.Database {
	if mongoInstance == nil {
		uri := os.Getenv("MONGODB_URI")
		if uri == "" {
			log.Fatalln("No mongo_uri connection string provided")
		}
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatalln("Connecting to mongodb_uri failed")
		}
		db := client.Database("gotwitter")
		mongoInstance = db
		return mongoInstance
	}
	return mongoInstance
}
