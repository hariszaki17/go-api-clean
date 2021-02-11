package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hariszaki17/go-api-clean/exception"
	"github.com/hariszaki17/go-api-clean/model"
	"github.com/hariszaki17/go-api-clean/service"
)

// UserController expose global
type UserController struct {
	UserService service.UserService
}
// NewUserController expose global
func NewUserController(userService *service.UserService) UserController {
	return UserController{
		UserService: *userService,
	}
}
// Route expose global
func (userController *UserController) Route(app *fiber.App) {
	app.Post("/api/users", userController.Create)
	app.Post("/api/users/login", userController.Login)
	app.Get("/api/users", userController.List)
	app.Delete("/api/users", userController.DeleteAll)
}

// Create expose global
func (userController *UserController) Create(c *fiber.Ctx) error {
	var request model.CreateUserRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response := userController.UserService.Create(request)
	return c.JSON(model.WebResponse{
		Code:	200,
		Status:	"OK",
		Data:	response,
	})
}

// List expose global
func (userController *UserController) List(c *fiber.Ctx) error{
	responses := userController.UserService.List()
	return c.JSON(model.WebResponse{
		Code:	200,
		Status:	"OK",
		Data:	responses,
	})
}

// DeleteAll expose global
func (userController *UserController) DeleteAll(c *fiber.Ctx) error{
	response := userController.UserService.DeleteAll()
	return c.JSON(model.WebResponse{
		Code:	200,
		Status:	"OK",
		Data: response,
	})
}

// Login expose global
func (userController *UserController) Login(c *fiber.Ctx) error {
	var request model.LoginUserRequest
	err := c.BodyParser(&request)
	
	exception.PanicIfNeeded(err)

	response := userController.UserService.Login(request)
	return c.JSON(model.WebResponse{
		Code: 200,
		Status: "OK",
		Data: response,
	})

}