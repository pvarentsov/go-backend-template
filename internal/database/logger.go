package database

import (
	"context"
	"fmt"
	"time"

	"go-backend-template/internal/util/contexts"

	"github.com/jackc/pgx/v4"
)

type Logger struct{}

func (l *Logger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	reqInfo, _ := contexts.GetReqInfo(ctx)

	if sql, ok := data["sql"]; ok {
		fmt.Printf("%s - [Database] TraceId: %s; UserId: %d; SQL: %s;\n\n",
			time.Now().Format(time.RFC1123),
			reqInfo.TraceId,
			reqInfo.UserId,
			sql,
		)
	}
}
