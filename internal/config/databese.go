package config

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create Connection
func NewMongo(config *viper.Viper, log *logrus.Logger) *mongo.Client {
	MongoDB := config.GetString("databese.mongo.url")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Fatalf("Failed to Connect Databese : %v",err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to Connect Databese : %v",err)
	}

	log.Info("Connected to MongoDB!")
	return client
}

// Client Database Instance
//var Client *mongo.Client = NewMongo()

// Open Collection

