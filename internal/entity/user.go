package entity

type SetUserInfoReq struct {
	UserId   int64  `json:"userId" binding:"required"`
	UserName string `json:"userName" binding:"min=3,max=64"`
	Gender   int    `json:"gender" binding:"oneof=1 2"`
	Age      int    `json:"age" binding:"min=1,max=100"`
}
