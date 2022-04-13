package main

import (
	"context"
	"log"

	"go-backend-template/api/cli"
	"go-backend-template/api/http"
	"go-backend-template/internal/database"
	"go-backend-template/internal/usecase"
	"go-backend-template/internal/util/crypto"
)

func main() {
	ctx := context.Background()
	parser := cli.NewParser()

	conf, err := parser.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbClient := database.NewClient(ctx, conf.Database())

	err = dbClient.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer dbClient.Close()

	crypto := crypto.NewCrypto()
	dbService := database.NewService(dbClient)
	usecases := usecase.NewUsecases(dbService, conf.Usecase(), crypto)

	server := http.NewServer(conf.HTTP(), crypto, &usecases)

	log.Fatal(server.Listen())
}
