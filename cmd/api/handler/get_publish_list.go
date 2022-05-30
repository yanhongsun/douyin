package handler

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/video"
	"douyin/pkg/errno"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPublishList(c *gin.Context) {

	var queryVar struct {
		Token  string `json:"token" form:"token"`
		UserId int64  `json:"user_id" form:"user_id"`
	}

	//TODO错误处理
	if err := c.BindQuery(&queryVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}

	//token处理

	if queryVar.Token != "" {
		userId, err := strconv.ParseInt(queryVar.Token, 10, 64)
		if err != nil {
			fmt.Println("handler.GetPublishList()->strconv.ParseInt error")
			return
		}
		//TODO错误处理
		if userId < 0 {
			SendResponse(c, errno.ParamErr, nil)
			return
		}
	}

	req := &video.GetPublishListRequest{Token: queryVar.Token, UserId: queryVar.UserId}

	videoList, err := rpc.GetPublishList(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, videoList)
}
