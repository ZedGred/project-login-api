package config

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create Connection
func NewMongo(config *viper.Viper, log *logrus.Logger) *mongo.Client {
	fmt.Println("aman")
	MongoDB := config.GetString("database.mongo.url")
	if MongoDB == "" {
		log.Fatal("MongoDB URL is not set")
	}

	fmt.Println(MongoDB)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	fmt.Println("aman")
	if err != nil {
		log.Errorf("Failed to Connect Databese : %v",err)
	}
	fmt.Println("aman")

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

