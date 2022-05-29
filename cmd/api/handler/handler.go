package handler

import (
	"douyin/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseFeed struct {
	Code     int32       `json:"status_code"`
	Message  string      `json:"status_msg"`
	NextTime int64       `json:"next_time"`
	Data     interface{} `json:"video_list"`
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

type Response struct {
	Code    int32       `json:"status_code"`
	Message string      `json:"status_msg"`
	Data    interface{} `json:"video_list"`
}

// SendResponse pack response
func SendResponse(c *gin.Context, err error, videolist interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    videolist,
	})
}
