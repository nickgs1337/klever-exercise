package cryptocurrency

import (
	"go.mongodb.org/mongo-driver/mongo"
	"klever/utils"
)

const collection = "cryptocurrency"

func getCollection() *mongo.Collection {
	return utils.GetDatabase().Collection(collection)
}
