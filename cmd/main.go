package main

import (
	"golangApiBoilerplate/internals/core/services"
	"golangApiBoilerplate/internals/handlers"
	"golangApiBoilerplate/internals/repositories"
	"golangApiBoilerplate/internals/server"
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
