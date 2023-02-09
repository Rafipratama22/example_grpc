package mongoconfig

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConfig struct {
	Client *mongo.Client
}

func SetupMongo() *MongoConfig {
	ctx := context.TODO()
	// Connect to MongoDB
	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017/")
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return &MongoConfig{
		Client: mongoclient,
	}
}
