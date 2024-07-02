package ports

import (
	"github.com/gofiber/fiber/v2"
	"goHexagonalBlog/internals/core/domain"
)

type IUserService interface {
	Get(email string) (domain.User, error)
	GetList() ([]domain.User, error)
}

type IUserRepository interface {
	Get(email string) (domain.User, error)
	GetList() ([]domain.User, error)
}

type IUserHandlers interface {
	Get(c *fiber.Ctx) error
	GetList(c *fiber.Ctx) error
}

type IServer interface {
	Initialize()
}
