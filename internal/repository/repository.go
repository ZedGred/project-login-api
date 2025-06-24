package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	DB *mongo.Collection
}

func (u *Repository) Create(ctx context.Context, entity interface{}) error {
	_, err := u.DB.InsertOne(ctx, entity)
	return err
}

func (u *Repository) Update(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	result, err := u.DB.UpdateOne(ctx, filter, bson.M{"$set": update})
	return result, err
}

func (u *Repository) CountByID(ctx context.Context, filter interface{}) (int64, error) {
	count, err := u.DB.CountDocuments(ctx, filter)
	return count, err
}
