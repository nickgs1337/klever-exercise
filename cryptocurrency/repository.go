package cryptocurrency

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"klever/utils"
)

const collection = "cryptocurrency"

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
