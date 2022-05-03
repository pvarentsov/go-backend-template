package contexts

import "context"

type reqInfoKey string

const (
	key reqInfoKey = "req-info"
)

type ReqInfo struct {
	UserId  int64
	TraceId string
}

func WithReqInfo(ctx context.Context, info ReqInfo) context.Context {
	return context.WithValue(ctx, key, info)
}

func GetReqInfo(ctx context.Context) (ReqInfo, bool) {
	reqInfo, ok := ctx.Value(key).(ReqInfo)
	if ok {
		return reqInfo, true
	}

	return ReqInfo{}, false
}
