package repository

import (
	"context"
	"project/internal/entity"
	"project/internal/utils"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	Repository
	Log *logrus.Logger
}

func NewUserRepository(db *mongo.Client, log *logrus.Logger) *UserRepository {
	collection := utils.OpenCollection(db, "user")
	return &UserRepository{
		Repository: Repository{DB: collection},
		Log:        log,
	}
}

func (u *UserRepository) FindByEmail(ctx context.Context, filter interface{}) (*entity.User, error) {
	var user entity.User
	err := u.DB.FindOne(ctx, bson.M{"email": filter}).Decode(&user)
	return &user, err
}

func (r *UserRepository) UpdateTokens(ctx context.Context, userId, token, refreshToken string) error {
	update := bson.M{
		"token":         token,
		"refresh_token": refreshToken,
		"updated_at":    time.Now().UTC(),
	}

	filter := bson.M{"_id": userId}
	upsert := true

	opts := options.UpdateOptions{Upsert: &upsert}

	_, err := r.DB.UpdateOne(ctx, filter, bson.M{"$set": update}, &opts)
	return err
}



