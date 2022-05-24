package handlers

import (
	"github.com/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserParam 用户登录/注册请求信息
type UserParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// QueryUserParam 查询用户请求信息
type QueryUserParam struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

// UserInfo 存储用户信息服务返回的用户信息
type UserInfo struct {
	UserID        int64  `json:"user_id"`
	Username      string `json:"username"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

// Response 响应信息
type Response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse 返回响应信息
func SendResponse(c *gin.Context, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}
