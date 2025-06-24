package usecase

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"project/internal/entity"
	"project/internal/model"
	"project/internal/model/converter"
	"project/internal/repository"
	"project/internal/utils"
)

type UserUseCase struct {
	DB             *mongo.Client
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository *repository.UserRepository
	Jwt            []byte
}

func NewUserUseCase(db *mongo.Client, logger *logrus.Logger, validate *validator.Validate,
	userRepository *repository.UserRepository, jwt []byte) *UserUseCase {
	return &UserUseCase{
		DB:             db,
		Log:            logger,
		Validate:       validate,
		UserRepository: userRepository,
		Jwt:            jwt,
	}
}

func (u *UserUseCase) Create(request *model.UserCreateRequest) (*model.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err := u.Validate.Struct(request)
	if err != nil {
		u.Log.Warnf("invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	count, err := u.UserRepository.CountByID(ctx, bson.M{"email": request.Email})
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if count > 0 {
		return nil, fiber.ErrConflict
	}

	password, err := utils.HashPassword(request.Password)
	if err != nil {
		u.Log.Warnf("failed generete bcrypt hash : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	user := &entity.User{
		ID:        primitive.NewObjectID(),
		Username:  request.Username,
		Email:     request.Email,
		Password:  password,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	AccesToken, RefreshToken, err := utils.GenereteAllToken(user, u.Jwt)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	user.Token = &AccesToken
	user.RefreshToken = &RefreshToken

	err = u.UserRepository.Create(ctx, user)
	if err != nil {
		u.Log.Errorf("failed to create user : %v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToResponse(user), nil
}

func (u *UserUseCase) Login(request *model.UserLoginRequest) (*model.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err := u.Validate.Struct(request)
	if err != nil {
		u.Log.Warnf("invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	user, err := u.UserRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		return nil, fiber.ErrNotFound
	}

	err = utils.VerifyPassword(user.Password, request.Password)
	if err != nil {
		return nil, fiber.ErrUnauthorized
	}

	AccesToken, RefreshToken, err := utils.GenereteAllToken(user, u.Jwt)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}
	user.Token = &AccesToken
	user.RefreshToken = &RefreshToken

	err = u.UserRepository.UpdateTokens(ctx, user.ID.Hex(), AccesToken, RefreshToken)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserTokenToResponse(user), nil
}
