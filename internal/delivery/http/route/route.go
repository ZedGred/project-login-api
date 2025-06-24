package route

import (
	"project/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct{
	App *fiber.App
	Midlleware fiber.Handler
	UserController *http.UserController
 }

 func(route *RouteConfig) Setup(){
	route.SetupAuthRoute()
 }

 func(route *RouteConfig) SetupAuthRoute(){
	route.App.Use(route.Midlleware)
	route.App.Post("/api/register",route.UserController.Register)
	route.App.Post("/api/login",route.UserController.Login)
 }