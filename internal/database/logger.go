package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type Logger struct{}

func (l *Logger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	if sql, ok := data["sql"]; ok {
		log.Printf("[Database] %s: %s\n", msg, sql)
	}
}
