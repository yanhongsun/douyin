package handlers

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RequestParam req format for register/login
type RequestParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserInfoParam req format for get_user_info
type UserInfoParam struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
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
