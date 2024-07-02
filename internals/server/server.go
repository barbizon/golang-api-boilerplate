package server

import (
	"github.com/gofiber/fiber/v2"
	"golangApiBoilerplate/internals/handlers"
	"golangApiBoilerplate/internals/middleware"
	"log"
)

type IServer interface {
	Initialize()
}
type Server struct {
	userHandlers handlers.IUserHandlers
}

func NewServer(uHandlers handlers.IUserHandlers) *Server {
	return &Server{
		userHandlers: uHandlers,
	}
}

func (s *Server) Initialize() {
	app := fiber.New()
	v1 := app.Group("/v1")

	userRoutes := v1.Group("/users", middleware.TestLoggerMiddleware, middleware.TestAuthenticationMiddleware)
	userRoutes.Get("/:email", s.userHandlers.Get)
	userRoutes.Get("", s.userHandlers.GetList)

	err := app.Listen(":5000")
	if err != nil {
		log.Fatal(err)
	}
}
