package http
//
//import (
//	"context"
//	"github.com/J-guanghua/client/middeware"
//	http2 "github.com/J-guanghua/client/http"
//	"net/http"
//	"time"
//)
//
//type Option func(*Client)
//type HandlerFunc func(r *http.Request,response *http2.Response) error
//type Unmarshal func(response *http2.Response,reply interface{}) error
//
//type Client struct {
//	http.Transport
//	ms      []middeware.Middleware
//	Decode  Unmarshal
//}
//
//func timeout(t time.Duration) Option  {
//	return func(client *Client) {
//		client.Timeout = t
//	}
//}
//func Middleware(ms ...middeware.Middleware) Option  {
//	return func(client *Client) {
//		client.ms = ms
//	}
//}
//func Decode(u Unmarshal) Option  {
//	return func(client *Client) {
//		client.Decode = u
//	}
//}
//func NewClient(ops ...Option) *Client {
//	c := &Client{
//		Client:&http.Client{},
//	}
//	for _,o:=range ops{
//		o(c)
//	}
//	return c
//}
//
//func (c *Client) RoundTrip(req *http.Request) (*http.Response,error) {
//	return c.Do(req)
//}
//
//func (c *Client) send(ctx context.Context,req *http.Request,ms ...middeware.Middleware) (*http2.Response,error) {
//	return middeware.Chain(append(ms,c.ms...)...)(func(ctx context.Context, req *http.Request) (*http2.Response, error) {
//		select {
//		case <-ctx.Done():
//			return nil,ctx.Err()
//		}
//		resp,err := c.RoundTrip(req)
//		return &http2.Response{Response:resp}, err
//	})(ctx,req)
//}
//
//func (c *Client)  HandlerFunc(req *http.Request,handler HandlerFunc,ms ...middeware.Middleware)  error {
//	resp,err := c.send(req.Context(),req,ms...)
//	if err != nil {
//		return err
//	}
//	return handler(req,resp)
//}
//
//func (c *Client)  ReadReply(req *http.Request,reply interface{},ms ...middeware.Middleware)  error {
//	resp,err := c.send(req.Context(),req,ms...)
//	if err != nil {
//		return err
//	}
//	return c.Decode(resp,&reply)
//}