package cache
import (
	"context"
	middeware2 "github.com/J-guanghua/client/middeware"
	http2 "github.com/J-guanghua/client/http"
	"net/http"
)

type cache map[string]interface{}

func NewCache() *cache {
	return new(cache)
}
func (c *cache) Get(ctx context.Context) *http2.Response {
	return nil
}
func (c *cache) Set(ctx context.Context,resp *http2.Response) error {
	return nil
}
type CacheInterface interface {
	Get(ctx context.Context)(*http2.Response)
	Set(ctx context.Context,resp *http2.Response) error
}

func HttpCache(c CacheInterface) middeware2.Middleware {

	return func(handler middeware2.Handler) middeware2.Handler {
		return func(ctx context.Context,req *http.Request) (*http2.Response,error) {
			if resp := c.Get(ctx);resp != nil{
				return resp, nil
			}
			resp,err := handler(ctx,req)
			if err != nil {
				return nil,err
			}
			c.Set(ctx,resp)
			return resp,err
		}
	}
}