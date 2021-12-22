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

func List() ([]Cryptocurrency, error) {
	cursor, err := getCollection().Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var results []Cryptocurrency
	for cursor.Next(context.TODO()) {
		var currency Cryptocurrency
		err := cursor.Decode(&currency)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, currency)
	}

	return results, nil
}

func GetBySymbol(symbol string) (*Cryptocurrency, error) {
	var cryptocurrency Cryptocurrency
	err := getCollection().FindOne(context.TODO(), bson.D{{"_id", symbol}}).Decode(&cryptocurrency)

	if err != nil {
		return nil, errors.New("currency not found")
	}

	return &cryptocurrency, err
}

func Update(symbol string, cryptocurrency *Cryptocurrency) (*Cryptocurrency, error) {
	_, err := GetBySymbol(cryptocurrency.Symbol)
	if err != nil {
		return nil, err
	}
	result, err := getCollection().UpdateOne(
		context.TODO(),
		bson.M{"_id": cryptocurrency.Symbol},
		bson.D{
			{"$set", bson.D{{"_id", cryptocurrency.Symbol}}},
			{"$set", bson.D{{"name", cryptocurrency.Name}}},
			{"$set", bson.D{{"icon_url", cryptocurrency.IconURL}}},
		},
	)
	if err != nil {
		return nil, errors.New("error while updating currency")
	}

	log.Printf("%d updated!", result.ModifiedCount)

	return cryptocurrency, nil
}

func Delete(symbol string) error {
	_, err := getCollection().DeleteOne(context.TODO(), bson.M{"_id": symbol})
	if err != nil {
		return errors.New("error while deleting")
	}
	return nil
}

func UpVote(symbol string) (*Cryptocurrency, error) {
	_, err := GetBySymbol(symbol)
	if err != nil {
		return nil, err
	}
	result, err := getCollection().UpdateOne(
		context.TODO(),
		bson.M{"_id": symbol},
		bson.D{
			{"$inc", bson.D{{"votes", 1}}},
		},
	)
	if err != nil {
		return nil, errors.New("error while updating currency")
	}
	log.Printf("%d updated!", result.ModifiedCount)

	return GetBySymbol(symbol)
}

func DownVote(symbol string) (*Cryptocurrency, error) {
	_, err := GetBySymbol(symbol)
	if err != nil {
		return nil, err
	}
	result, err := getCollection().UpdateOne(
		context.TODO(),
		bson.M{"_id": symbol},
		bson.D{
			{"$inc", bson.D{{"votes", -1}}},
		},
	)
	if err != nil {
		return nil, errors.New("error while updating currency")
	}
	log.Printf("%d updated!", result.ModifiedCount)

	return GetBySymbol(symbol)
}

func getCollection() *mongo.Collection {
	return utils.GetDatabase().Collection(collection)
}
