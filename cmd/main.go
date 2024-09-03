package main

import (
	"toktok-backend-v1.0.1/internal/adapter/persistence/mysql"
	"toktok-backend-v1.0.1/internal/adapter/persistence/mysql/repository"
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

	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)

}
