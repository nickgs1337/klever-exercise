package cryptocurrency

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"klever/utils"
	"log"
)

const collection = "cryptocurrency"

func Create(cryptocurrency *Cryptocurrency) (*Cryptocurrency, error) {
	_, err := GetBySymbol(cryptocurrency.Symbol)
	if err == nil {
		return nil, errors.New("currency already registered")
	}

	result, err := getCollection().InsertOne(context.TODO(), cryptocurrency)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s created!", result.InsertedID)

	return cryptocurrency, nil
}

func GetBySymbol(symbol string) (*Cryptocurrency, error) {
	var cryptocurrency Cryptocurrency
	err := getCollection().FindOne(context.TODO(), bson.D{{"_id", symbol}}).Decode(&cryptocurrency)

	if err != nil {
		return nil, errors.New("currency not found")
	}

	return &cryptocurrency, err
}

func getCollection() *mongo.Collection {
	return utils.GetDatabase().Collection(collection)
}
