package handlers

import (
	"github.com/gofiber/fiber/v2"
	"golangApiBoilerplate/internals/core/services"
	"log"
	"time"
)

type UserHandlers struct {
	userService services.IUserService
}
type IUserHandlers interface {
	Get(c *fiber.Ctx) error
	GetList(c *fiber.Ctx) error
}

var _ IUserHandlers = (*UserHandlers)(nil)

func NewUserHandlers(userService services.IUserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

func (h *UserHandlers) Get(c *fiber.Ctx) error {
	start := time.Now()
	log.Printf("User Get: Method: %s, Path: %s, Status: %d, Duration: %s",
		c.Method(), c.Path(), c.Response().StatusCode(), time.Since(start))

	var email string

	data, err := h.userService.Get(email)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(data)
}

func (h *UserHandlers) GetList(c *fiber.Ctx) error {
	start := time.Now()
	log.Printf("User GetList: Method: %s, Path: %s, Status: %d, Duration: %s",
		c.Method(), c.Path(), c.Response().StatusCode(), time.Since(start))
	data, err := h.userService.GetList()
	if err != nil {
		return err
	}
	return c.Status(200).JSON(data)
}
