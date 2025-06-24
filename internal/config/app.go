package config

import (
	"project/internal/delivery/http/middleware"
	"project/internal/delivery/http/route"
	"project/internal/repository"
	"project/internal/usecase"
	"project/internal/delivery/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type BootstrapConfig struct {
	DB *mongo.Client
	App *fiber.App
	Config *viper.Viper
	Log *logrus.Logger
	Validate *validator.Validate
	Jwt []byte
}

func Bootstrap(config *BootstrapConfig) {
	//Setup Repository
	userRepository := repository.NewUserRepository(config.DB,config.Log)
	
	//Setup Usecase 
	userUseCase := usecase.NewUserUseCase(config.DB,config.Log,config.Validate,userRepository,config.Jwt)

	//Setup Controller
	UserController := http.NewUserController(userUseCase,config.Log)

	//Setup Midlleware
	authMiddleware := middleware.AuthMiddleware(config.Jwt)

	Route := route.RouteConfig{
		App: config.App,
		Midlleware: authMiddleware,
		UserController: UserController,
		}

	Route.Setup()
}
	