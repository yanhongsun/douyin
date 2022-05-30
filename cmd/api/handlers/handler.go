package handlers

import (
	"douyin/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int32       `json:"status_code"`
	Message string      `json:"status_msg"`
	Data    interface{} `json:"user_list"`
}
type BaseResponse struct {
	Code    int32  `json:"status_code"`
	Message string `json:"status_msg"`
}

// SendResponse pack response
func SendResponse(c *gin.Context, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}

func SendBaseResponse(c *gin.Context, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, BaseResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
	})
}
