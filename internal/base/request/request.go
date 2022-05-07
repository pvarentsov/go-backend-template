package request

import "context"

type requestInfoKey int

const (
	key requestInfoKey = iota
)

type RequestInfo struct {
	UserId  int64
	TraceId string
}

func WithRequestInfo(ctx context.Context, info RequestInfo) context.Context {
	return context.WithValue(ctx, key, info)
}

func GetRequestInfo(ctx context.Context) (requestInfo RequestInfo, ok bool) {
	requestInfo, ok = ctx.Value(key).(RequestInfo)
	return
}
