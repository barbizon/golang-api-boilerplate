package main

import (
	"goHexagonalBlog/internals/core/services"
	"goHexagonalBlog/internals/handlers"
	"goHexagonalBlog/internals/repositories"
	"goHexagonalBlog/internals/server"
)

func main() {
	userHandlers := userHandlersDI()
	//server
	s := server.NewServer(userHandlers)
	s.Initialize()
}

func userHandlersDI() *handlers.UserHandlers {
	//repositories
	userRepository, _ := repositories.NewUserRepository()
	//services
	userService := services.NewUserService(userRepository)
	//handlers
	userHandlers := handlers.NewUserHandlers(userService)

	return userHandlers
}
