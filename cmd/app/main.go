package main

import (
	"log"

	"toktok-backend-v1.0.1/internal/adapter/authentication/token"
	"toktok-backend-v1.0.1/internal/adapter/persistence/mysql"
	"toktok-backend-v1.0.1/internal/adapter/persistence/mysql/repository"
	"toktok-backend-v1.0.1/internal/adapter/presentation/handler"
	"toktok-backend-v1.0.1/internal/adapter/presentation/middleware"
	"toktok-backend-v1.0.1/internal/adapter/presentation/router"
	"toktok-backend-v1.0.1/internal/config"
	"toktok-backend-v1.0.1/internal/core/service"
)

func main() {
	config, err := config.New(".toml")
	if err != nil {
		panic(err)
	}

	database, err := mysql.NewDatabase(config)
	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate()
	if err != nil {
		panic(err)
	}

	tokenService, err := token.NewTokenService(config)
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	authService := service.NewAuthService(userRepository, tokenService)
	authHandler := handler.NewAuthHandler(authService)

	handlerSet := router.HandlerSet{
		UserHandler: userHandler,
		AuthHandler: authHandler,
	}

	guardMiddleware := middleware.NewGuardMiddlware(tokenService)

	middlewareSet := router.MiddlewareSet{
		GuardMiddleware: guardMiddleware,
	}

	router := router.NewRouter(config, handlerSet, middlewareSet)

	log.Fatal(router.Listen())

}
