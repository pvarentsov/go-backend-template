package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-backend-template/internal/util/contexts"
)

type reqInfoKey = string

const (
	key reqInfoKey = "req-info"
)

func SetTraceId(c *gin.Context, traceId string) {
	info, exists := c.Get(key)
	if exists {
		parsedInfo := info.(contexts.ReqInfo)
		parsedInfo.TraceId = traceId

		c.Set(key, parsedInfo)

		return
	}

	c.Set(key, contexts.ReqInfo{TraceId: traceId})
}

func SetUserId(c *gin.Context, userId int64) {
	info, exists := c.Get(key)
	if exists {
		parsedInfo := info.(contexts.ReqInfo)
		parsedInfo.UserId = userId

		c.Set(key, parsedInfo)

		return
	}

	c.Set(key, contexts.ReqInfo{UserId: userId})
}

func WithReqInfo(c *gin.Context) context.Context {
	info, ok := c.Get(key)
	if ok {
		return contexts.WithReqInfo(c, info.(contexts.ReqInfo))
	}

	return contexts.WithReqInfo(c, contexts.ReqInfo{})
}
