package handlers

import (
	"douyin/pkg/errno"
	"net/http"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

// RequestParam req format for register/login
type RequestParam struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// UserInfoParam req format for get_user_info
type UserInfoParam struct {
	UserID int64  `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
}

// UserInfo user info format
type UserInfo struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

// Response resp format of register/login
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserID     int64  `json:"user_id"`
	Token      string `json:"token"`
}

// UserInfoResponse resp format of get_user_info
type UserInfoResponse struct {
	StatusCode int32    `json:"status_code"`
	StatusMsg  string   `json:"status_msg"`
	Data       UserInfo `json:"user"`
}
type ResponseV struct {
	Code    int32       `json:"status_code"`
	Message string      `json:"status_msg"`
	Data    interface{} `json:"video_list"`
}

type ResponseFeed struct {
	Code     int32       `json:"status_code"`
	Message  string      `json:"status_msg"`
	NextTime int64       `json:"next_time"`
	Data     interface{} `json:"video_list"`
}

// SendResponse send response of register/login
func SendResponse(c *gin.Context, err error, userID int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		UserID:     userID,
		Token:      token,
	})
}

// SendUserInfoResponse send response of get_user_info
func SendUserInfoResponse(c *gin.Context, err error, userInfo *UserInfo) {
	Err := errno.ConvertErr(err)
	klog.Info("================================")
	klog.Info(userInfo)
	klog.Info(Err)
	klog.Info("================================")
	c.JSON(http.StatusOK, UserInfoResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		Data: UserInfo{
			ID:            userInfo.ID,
			Name:          userInfo.Name,
			FollowCount:   userInfo.FollowerCount,
			FollowerCount: userInfo.FollowerCount,
			IsFollow:      userInfo.IsFollow,
		},
	})
}

// SendResponse pack response
func SendResponseFeed(c *gin.Context, err error, videolist interface{}, nexttime int64) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, ResponseFeed{
		Code:     Err.ErrCode,
		Message:  Err.ErrMsg,
		NextTime: nexttime,
		Data:     videolist,
	})
}

// SendResponse pack response
func SendResponseV(c *gin.Context, err error, videolist interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, ResponseV{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    videolist,
	})
}
