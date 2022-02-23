package log

import "context"

const (
	ContextKeyRequestId = "ctx-key::request-id"
)

func WithRequestId(ctx context.Context, rId string) context.Context {
	return context.WithValue(ctx, ContextKeyRequestId, rId)
}

func GetRequestId(ctx context.Context) string {
	if rIdObj := ctx.Value(ContextKeyRequestId); rIdObj != nil {
		return rIdObj.(string)
	}
	return ""
}
