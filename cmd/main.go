package main

import (
	"toktok-backend-v1.0.1/internal/adapter/persistence/mysql"
	"toktok-backend-v1.0.1/internal/config"
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
}
