# Go Backend Template

## Examples

Database (pgx + goqu):
  * [Client](./internal/database/client.go)
  * [Service](./internal/database/service.go)
  * [Transaction](./internal/database/transaction.go)
  * [Repository](./internal/database/repository/user.go)

Usecases:
  * [Auth](./internal/usecase/auth.go)
  * [User](./internal/usecase/user.go)
  * [Usecase with transaction](./internal/usecase/transaction.go)

HTTP Server (gin):
  * [Server](./api/http/server.go)
  * [Router](./api/http/router.go)
  * [Main file](./cmd/http/main.go) 

Request Collection:
  * [InsomniaV4](./assets/api-collection.insomnia-v4.json)

## Makefile

```shell
$ make

Usage: make [command]

Commands:
 build-http                    Build http server

 migration-create name={name}  Create migration
 migration-up                  Up migrations
 migration-down                Down last migration

 docker-up                     Up docker services
 docker-down                   Down docker services

 fmt                           Format source code

Requirements:
 migrate                       Migration tool: https://github.com/golang-migrate/migrate
```
## License

This project is licensed under the [MIT License](https://github.com/pvarentsov/go-backend-template/blob/main/LICENSE).
