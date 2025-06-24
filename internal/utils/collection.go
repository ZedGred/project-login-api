package utils

import(
	"go.mongodb.org/mongo-driver/mongo"
)

func OpenCollection(client *mongo.Client, collectioname string) *mongo.Collection {
	var collections *mongo.Collection = client.Database("cluster0").Collection(collectioname)

	return collections
}

