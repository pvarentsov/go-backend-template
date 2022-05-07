package main

import (
	"context"
	"log"

	"go-backend-template/api/cli"
	"go-backend-template/api/http"

	authImpl "go-backend-template/internal/auth/impl"
	cryptoImpl "go-backend-template/internal/base/crypto/impl"
	databaseImpl "go-backend-template/internal/base/database/impl"
	userImpl "go-backend-template/internal/user/impl"
)

func main() {
	ctx := context.Background()
	parser := cli.NewParser()

	conf, err := parser.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbClient := databaseImpl.NewClient(ctx, conf.Database())

	err = dbClient.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer dbClient.Close()

	crypto := cryptoImpl.NewCrypto()
	dbService := databaseImpl.NewService(dbClient)

	userRepositoryOpts := userImpl.UserRepositoryOpts{
		ConnManager: dbService,
	}
	userRepository := userImpl.NewUserRepository(userRepositoryOpts)

	authServiceOpts := authImpl.AuthServiceOpts{
		Crypto:         crypto,
		Config:         conf.Auth(),
		UserRepository: userRepository,
	}
	authService := authImpl.NewAuthService(authServiceOpts)

	userUsecasesOpts := userImpl.UserUsecasesOpts{
		TxManager:      dbService,
		UserRepository: userRepository,
		Crypto:         crypto,
	}
	userUsecases := userImpl.NewUserUsecases(userUsecasesOpts)

	serverOpts := http.ServerOpts{
		UserUsecases: userUsecases,
		AuthService:  authService,
		Crypto:       crypto,
		Config:       conf.HTTP(),
	}
	server := http.NewServer(serverOpts)

	log.Fatal(server.Listen())
}
