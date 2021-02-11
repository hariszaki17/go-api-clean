package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hariszaki17/go-api-clean/config"
	"github.com/hariszaki17/go-api-clean/controller"
	"github.com/hariszaki17/go-api-clean/exception"
	"github.com/hariszaki17/go-api-clean/repository"
	"github.com/hariszaki17/go-api-clean/service"
)

func main()  {
	configuration := config.New()
	database := config.NewMongoDatabase(configuration)

	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(&userRepository)
	userController := controller.NewUserController(&userService)
	
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	userController.Route(app)

	err := app.Listen(":3003")
	exception.PanicIfNeeded(err)
}