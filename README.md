# Go Backend Template

## Makefile

```shell
$ make

Usage: make [command]

Commands:
   build-http                                Build http server

   migration-create name={migration_name}    Create migration
   migration-up                              Up migrations
   migration-down                            Down last migration

   docker-up                                 Up docker services
   docker-down                               Down docker services

   fmt                                       Format source code
   
Requirements:
   migrate                                   Migration tool: https://github.com/golang-migrate/migrate
```
