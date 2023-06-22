package auth

import (
	"context"
	"fmt"
	http2 "github.com/J-guanghua/client/http"
	"github.com/J-guanghua/client/middeware"
	"net/http"
)

type Accout struct {
	addr string
	host string
	err error
	authorization string
}

func NewAuth() *Accout {
	return &Accout{
		err:fmt.Errorf("授权异常:%s","未初始化账户信息"),
	}
}
func (a *Accout) Error() error {
	return nil
}

func (a *Accout) AuthFunc(r *http.Request) error {
	r.Host = a.host
	r.Header.Set("authorization",a.authorization)
	return nil
}

type AuthInterface interface {
	AuthFunc(r *http.Request) error
	Error() error
}

func Authorization(auth AuthInterface) middeware.Middleware {

	return func(handler middeware.Handler) middeware.Handler {
		return func(ctx context.Context,req *http.Request) (*http2.Response,error) {
			if err := auth.Error();err != nil{
				return nil,err
			}
			err := auth.AuthFunc(req)
			if err != nil {
				return nil,err
			}
			return handler(ctx,req)
		}
	}
}
