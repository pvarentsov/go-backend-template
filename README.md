# Go Backend Template

## Examples

<br>
<table align="center">
<thead>
<tr>
<th>Layer</th>
<th>Implementations</th>
</tr>
</thead>
<tbody>
<tr>
<td>Database (<a href="https://github.com/jackc/pgx">pgx</a> + <a href="https://github.com/doug-martin/goqu">goqu</a>)</td>
<td><a href="./internal/database/client.go">Client</a>, <a href="./internal/database/service.go">Service</a>, <a href="./internal/database/transaction.go">Transaction</a>, <a href="./internal/database/repository/user.go">User Repository</a></td>
</tr>
<tr>
<td>Model (<a href="https://github.com/go-ozzo/ozzo-validation">ozzo-validation</a>)</td>
<td><a href="./internal/model/user.go">User</a></td>
</tr>
<tr>
<td>Usecases</td>
<td><a href="./internal/usecase/auth.go">Auth</a>, <a href="./internal/usecase/user.go">User</a>, <a href="./internal/usecase/transaction.go">Usecase with transaction</a></td>
</tr>
<tr>
<td>HTTP Server (<a href="https://github.com/gin-gonic/gin">gin</a>)</td>
<td><a href="./api/http/server.go">Server</a>, <a href="./api/http/router.go">Router</a>, <a href="./cmd/http/main.go">Main file</a></td>
</tr>
</tbody>
</table>
<br>

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

## Request Collection
* [InsomniaV4](./assets/api-collection.insomnia-v4.json)

## License

This project is licensed under the [MIT License](https://github.com/pvarentsov/go-backend-template/blob/main/LICENSE).
