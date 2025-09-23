package entity

import "time"

type SetUserInfoReq struct {
	UserId   int64  `json:"userId" binding:"required"`
	UserName string `json:"userName" binding:"min=3,max=64"`
	Gender   int    `json:"gender" binding:"oneof=1 2"`
	Age      int    `json:"age" binding:"min=1,max=100"`
}
type LoginReq struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResp struct {
	UserId int    `json:"userId"`
	Token  string `json:"token"`
}

type UserListResp struct {
	Users []User `json:"users"`
}

type User struct {
	Id        int       `json:"id"`
	NickName  string    `json:"nick_name"`
	Account   string    `json:"account"`
	Password  string    `json:"password"`
	Icon      string    `json:"icon"`
	Gender    int       `json:"gender"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
