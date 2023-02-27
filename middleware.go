package client


import (
	"context"
	"net/http"
)


type Handler func(ctx context.Context, req *http.Request) (*Response, error)

type Middleware func(Handler) Handler

func Chain(m ...Middleware) Middleware {
	return func(next Handler) Handler {
		for i := len(m) - 1; i >= 0; i-- {
			next = m[i](next)
		}
		return next
	}
}

