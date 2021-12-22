package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func getConnection() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), getClientOptions())
	if err != nil {
		log.Fatal(err)
	}

	// validating the connection...
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetDatabase() *mongo.Database {
	return getConnection().Database(GetConfig().MongoDatabase)
}

func getClientOptions() *options.ClientOptions {
	credential := options.Credential{
		Username: GetConfig().MongoUser,
		Password: GetConfig().MongoPassword,
	}
	return options.Client().ApplyURI(GetConfig().MongoURI).SetAuth(credential)
}
