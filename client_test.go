package client


type UsersReq struct {
	Openid string
}
type UserReply struct {
	Nickname string `json:"nickname"`
	Openid string `json:"openid"`
	ImageUrl string `json:"image_url"`
	Tags []string `json:"tags"`
}