package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hariszaki17/go-api-clean/exception"
)

// NewFiberConfig expose global
func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}