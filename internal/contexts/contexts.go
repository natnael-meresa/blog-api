package contexts

import "context"

type userIDCtx struct{}

func FromUserID(ctx context.Context) uint64 {
	v := ctx.Value(userIDCtx{})
	if v != nil {
		if s, ok := v.(uint64); ok {
			return s
		}
	}
	return 0
}
