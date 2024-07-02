package middleware

import (
	"github.com/gofiber/fiber/v2"
	"goHexagonalBlog/internals/core/services"
	"time"
)

func TestLoggerMiddleware(c *fiber.Ctx) error {
	// Process the request
	err := c.Next()
	// Log the details
	services.PrintTimestamp(time.Now(), c.Method(), c.Path(), c.Response().StatusCode())
	// Return
	return err
}

func TestAuthenticationMiddleware(c *fiber.Ctx) error {
	// Process the request
	err := c.Next()
	// Log the details
	services.PrintTimestamp(time.Now(), c.Method(), c.Path(), c.Response().StatusCode())
	// Return
	return err
}
