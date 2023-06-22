package client

import (
	"context"
	http2 "github.com/J-guanghua/client/http"
	"github.com/J-guanghua/client/middeware/auth"
	"github.com/J-guanghua/client/middeware/cache"
	"net/http"
	"testing"
)


type UsersReq struct {
	Openid string
}
type UserReply struct {
	Nickname string `json:"nickname"`
	Openid string `json:"openid"`
	ImageUrl string `json:"image_url"`
	Tags []string `json:"tags"`
}
type WechatCliect struct {
	c *Client
}
func (wc *WechatCliect) GetUsers(ctx context.Context,req *UsersReq) (*UserReply,error) {
	r ,err := http2.NewResquest(http.MethodGet,"/api/users/opneid/{openid}",req)
	if err != nil {
		return nil,err
	}
	users := &UserReply{}
	return users, wc.c.ReadReply(r.WithContext(ctx),users)
}

var wc = WechatCliect{NewClient(Middleware(cache.HttpCache(cache.NewCache()),auth.Authorization(auth.NewAuth())))}

func TestRequest(t *testing.T)  {
	_,err := wc.GetUsers(context.Background(),&UsersReq{
		Openid: "dwdwJKkndjwwkndkjw-JHGp",
	})
	if err != nil {
		t.Error(err)
	}
}