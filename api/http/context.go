package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-backend-template/internal/util/contexts"
)

type reqInfoKeyType = string

const (
	reqInfoKey reqInfoKeyType = "req-info"
)

func setTraceId(c *gin.Context, traceId string) {
	info, exists := c.Get(reqInfoKey)
	if exists {
		parsedInfo := info.(contexts.ReqInfo)
		parsedInfo.TraceId = traceId

		c.Set(reqInfoKey, parsedInfo)

		return
	}

	c.Set(reqInfoKey, contexts.ReqInfo{TraceId: traceId})
}

func setUserId(c *gin.Context, userId int64) {
	info, exists := c.Get(reqInfoKey)
	if exists {
		parsedInfo := info.(contexts.ReqInfo)
		parsedInfo.UserId = userId

		c.Set(reqInfoKey, parsedInfo)

		return
	}

	c.Set(reqInfoKey, contexts.ReqInfo{UserId: userId})
}

func getReqInfo(c *gin.Context) contexts.ReqInfo {
	info, ok := c.Get(reqInfoKey)
	if ok {
		return info.(contexts.ReqInfo)
	}

	return contexts.ReqInfo{}
}

func contextWithReqInfo(c *gin.Context) context.Context {
	info, ok := c.Get(reqInfoKey)
	if ok {
		return contexts.WithReqInfo(c, info.(contexts.ReqInfo))
	}

	return contexts.WithReqInfo(c, contexts.ReqInfo{})
}
