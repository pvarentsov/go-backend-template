<h1 align="center">
    <img height="80" src="./assets/gopher-icon.gif" alt="Go"><br>Backend Template
</h1>

> Clean architecture based backend template in **Go**.

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
<td align="center">internal/util</td>
<td align="center">crypto (<a href="https://pkg.go.dev/golang.org/x/crypto/bcrypt">bcrypt</a> + <a href="https://github.com/golang-jwt/jwt">jwt</a> + <a href="https://github.com/gofrs/uuid">uuid</a>);<br>contexts; errors</td>
<td align="center">
    crypto: <a href="./internal/util/crypto/crypto.go">Crypto</a>;<br>
    contexts: <a href="./internal/util/contexts/context.go">Context</a>;
    errors: <a href="./internal/util/errors/error.go">Error</a>, <a href="./internal/util/errors/status.go">Status</a>
</td>
</tr>

<tr>
<td align="center">internal</td>
<td align="center">model (<a href="https://github.com/go-ozzo/ozzo-validation">ozzo-validation</a>)</td>
<td align="center"><a href="./internal/model/user.go">User</a></td>
</tr>

<tr>
<td align="center">internal</td>
<td align="center">database (<a href="https://github.com/jackc/pgx">pgx</a> + <a href="https://github.com/doug-martin/goqu">goqu</a>)</td>
<td align="center">
    <a href="./internal/database/db_client.go">Client</a>, 
    <a href="./internal/database/db_service.go">Service</a>,  
    <a href="./internal/database/repo_user.go">UserRepo</a>
</td>
</tr>

<tr>
<td align="center">internal</td>
<td align="center">usecase</td>
<td align="center">
    <a href="./internal/usecase/usecase_auth.go">Auth</a>, 
    <a href="./internal/usecase/usecase_user.go">User</a>
</td>
</tr>

<tr>
<td align="center">config</td>
<td align="center">config (<a href="https://github.com/kelseyhightower/envconfig">envconfig</a> + <a href="https://github.com/subosito/gotenv">gotenv</a>)</td>
<td align="center">
    <a href="./config/config.go">Config</a>,
    <a href="./config/env">ENV files</a>
</td>
</tr>

<tr>
<td align="center">api</td>
<td align="center">http (<a href="https://github.com/gin-gonic/gin">gin</a>)</td>
<td align="center">
    <a href="./api/http/server.go">Server</a>, 
    <a href="./api/http/router.go">Router</a>
</td>
</tr>

<tr>
<td align="center">api</td>
<td align="center">cli (<a href="https://github.com/alecthomas/kong">kong</a>)</td>
<td align="center">
    <a href="./api/cli/cli.go">Parser</a>
</td>
</tr>

<tr>
<td align="center">cmd</td>
<td align="center">http/main</td>
<td align="center"><a href="./cmd/http/main.go">Main file</a></td>
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
 test                          Run unit tests

Requirements:
 docker-compose                Docker Compose CLI: https://docs.docker.com/compose/reference
 migrate                       Migration CLI tool: https://github.com/golang-migrate/migrate

```

## HTTP Server

```shell
$ ./bin/http-server --help

Usage: http-server

Flags:
  -h, --help               Show context-sensitive help.
      --env-path=STRING    Path to env config file
```

**Configuration** is based on the environment variables. See [.env.template](./config/env/.env.template).

```shell
# Expose env vars before and start server
$ ./bin/http-server

# Expose env vars from the file and start server
$ ./bin/http-server --env-path ./config/env/.env
```

## Request Collection
* [InsomniaV4](./assets/api-collection.insomnia-v4.json)

## License

This project is licensed under the [MIT License](https://github.com/pvarentsov/go-backend-template/blob/main/LICENSE).
