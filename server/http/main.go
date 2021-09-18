package main

import (
	"log"
	"seafarer-backend/api"
	"seafarer-backend/domain"
	"seafarer-backend/server/http/router"
)

func main() {

	// Load Configuration
	config, err := domain.LoadConfiguration()
	if err != nil {
		log.Fatal("Error while load configuration, ", err.Error())
	}
	defer config.PostgresConnection.Close()
	defer config.Redis.Close()

	// Insert Handler Contract
	handler := api.NewHandler(&api.Contract{
		App:       config.App,
		Validator: config.Validator,
		Postgres:  config.Postgres,
		Mail:      config.Mail,
		Redis:     config.Redis,
	})

	// Register routes
	router.NewRouter(handler).RegisterRoutes()

	// Listening Http
	if err = domain.HttpListen(config.App); err != nil {
		log.Fatal("Error while listening http protocol, ", err.Error())
	}

}
