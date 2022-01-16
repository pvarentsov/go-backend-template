package main

import (
	"context"
	"log"

	"go-backend-template/api/cli"
	"go-backend-template/api/http"
	"go-backend-template/internal/database"
	"go-backend-template/internal/usecase"
)

func main() {
	ctx := context.Background()
	cliScheme := cli.NewCLI()

	conf, err := cliScheme.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbClient := database.NewClient(ctx, conf.Database())

	err = dbClient.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer dbClient.Close()

	dbService := database.NewService(dbClient)
	usecases := usecase.NewUsecases(dbService, conf.Usecase())

	server := http.NewServer(conf.HTTP(), &usecases)

	log.Fatal(server.Listen())
}
