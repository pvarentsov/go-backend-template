<h1 align="center"><img height="20" src="./assets/go-icon.png"> Go Backend Template</h1>

## Structure

<br>
<table align="center">
<thead>
<tr>
<th>Layer</th>
<th>Package</th>
<th>Implementations</th>
</tr>
</thead>
<tbody>
<tr>
<td align="center">internal</td>
<td align="center">database (<a href="https://github.com/jackc/pgx">pgx</a> + <a href="https://github.com/doug-martin/goqu">goqu</a>)</td>
<td>
    <a href="./internal/database/client.go">Client</a>, 
    <a href="./internal/database/service.go">Service</a>, 
    <a href="./internal/database/transaction.go">Transaction</a>, 
    <a href="./internal/database/repository/user.go">UserRepository</a>
</td>
</tr>
<tr>
<td align="center">internal</td>
<td align="center">model (<a href="https://github.com/go-ozzo/ozzo-validation">ozzo-validation</a>)</td>
<td><a href="./internal/model/user.go">User</a></td>
</tr>
<tr>
<td>internal</td>
<td align="center">dto</td>
<td>
    <a href="./internal/dto/add_user.go">AddUser</a>,
    <a href="./internal/dto/user.go">User</a>,
    <a href="./internal/dto/login_user.go">LoginUser</a>,
    <a href="./internal/dto/logged_user.go">LoggedUser</a>,
    <a href="./internal/dto/">...</a>
</td>
</tr>
<tr>
<td align="center">internal</td>
<td align="center">usecase</td>
<td>
    <a href="./internal/usecase/auth.go">Auth</a>, 
    <a href="./internal/usecase/user.go">User</a>, 
    <a href="./internal/usecase/transaction.go">Transaction</a> (usecase example with transaction)
</td>
</tr>
<tr>
<td align="center">api</td>
<td align="center">http (<a href="https://github.com/gin-gonic/gin">gin</a>)</td>
<td>
    <a href="./api/http/server.go">Server</a>, 
    <a href="./api/http/router.go">Router</a>
</td>
</tr>
<tr>
<td align="center">config</td>
<td align="center">config</td>
<td><a href="./config/config.go">Config</a></td>
</tr>
<tr>
<td align="center">cmd</td>
<td align="center">http/main</td>
<td><a href="./cmd/http/main.go">Main file</a></td>
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
 docker-compose                Docker Compose CLI: https://docs.docker.com/compose/reference
 migrate                       Migration CLI tool: https://github.com/golang-migrate/migrate

```

## HTTP Server

### Help

```shell
$ ./bin/http-server --help

Usage: http-server

Flags:
  -h, --help               Show context-sensitive help.
      --env-path=STRING    Path to env config file
```

### Configuration

Configuration is based on the environment variables. See [.env.template](./config/env/.env.template).

```shell
# Expose env vars before and start server
$ ./bin/http-server

# Expose env vars from the file and start server
$ ./bin/http-server --env-path ./config/env/.env
```

<a href="https://www.flaticon.com/free-icons/go-lang" title="go lang icons">Go lang icons created by Freepik - Flaticon</a>

## Request Collection
* [InsomniaV4](./assets/api-collection.insomnia-v4.json)

## License

This project is licensed under the [MIT License](https://github.com/pvarentsov/go-backend-template/blob/main/LICENSE).
