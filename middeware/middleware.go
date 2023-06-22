package middeware


import (
	"context"
	http2 "github.com/J-guanghua/client/http"
	"net/http"
)


type Handler func(ctx context.Context, req *http.Request) (*http2.Response, error)

type Middleware func(Handler) Handler

func Chain(m ...Middleware) Middleware {
	return func(next Handler) Handler {
		for i := len(m) - 1; i >= 0; i-- {
			next = m[i](next)
		}
		return next
	}
}

