package client

import (
	"context"
	"net/http"
	"time"
)

type Response struct {
	body []byte
	*http.Response
}
type Option func(*Client)
type HandlerFunc func(r *http.Request,response *Response) error
type UnCode func(response *Response,reply interface{}) error
type Client struct {
	*http.Client
	ms []Middleware
	timeout time.Duration
	Decode UnCode
}

func NewClient(ops ...Option) *Client {
	return &Client{}
}

func (c *Client) RoundTrip(req *http.Request) (*http.Response,error) {
	return c.Do(req)
}

func (c *Client) send(ctx context.Context,req *http.Request,ms ...Middleware) (*Response,error) {
	return Chain(append(ms,c.ms...)...)(func(ctx context.Context, req *http.Request) (*Response, error) {
		select {
		case <-ctx.Done():
			return nil,ctx.Err()
		}
		resp,err := c.RoundTrip(req)
		return &Response{Response:resp}, err
	})(ctx,req)
}

func (c *Client)  HandlerFunc(req *http.Request,handler HandlerFunc,ms ...Middleware)  error {
	resp,err := c.send(req.Context(),req,ms...)
	if err != nil {
		return err
	}
	return handler(req,resp)
}

func (c *Client)  ReadReply(req *http.Request,reply interface{},ms ...Middleware)  error {
	resp,err := c.send(req.Context(),req,ms...)
	if err != nil {
		return err
	}
	return c.Decode(resp,&reply)
}