package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sirupsen/logrus"

	"project/internal/model"
	"project/internal/usecase"
)

type UserController struct {
	Log         *logrus.Logger
	UserUseCase *usecase.UserUseCase
}

func NewUserController(usecase *usecase.UserUseCase, log *logrus.Logger) *UserController {
	return &UserController{
		Log:         log,
		UserUseCase: usecase,
	}
}

func (u *UserController) Register(c *fiber.Ctx) error {
	request := new(model.UserCreateRequest)
	err := c.BodyParser(request)
	if err != nil {
		return fiber.ErrBadRequest
	}

	response, err := u.UserUseCase.Create(request)
	if err != nil {
        log.Warnf("failed create acoount : %+v",err)
		return err
	}

	return c.Status(http.StatusCreated).JSON(model.SuccessRespones{
		Success: true,
		Message: "User created successfully",
		Data:    response})
}

func (u *UserController) Login(c *fiber.Ctx) error {
	request := new(model.UserLoginRequest)
	err := c.BodyParser(request)
	if err != nil {
		return fiber.ErrBadRequest
	}

	response, err := u.UserUseCase.Login(request)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(model.SuccessRespones{
		Success: true,
		Message: "Login succesfully",
		Data:    response,
	})

}
