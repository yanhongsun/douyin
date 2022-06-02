package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/video"
	"douyin/pkg/errno"

	"github.com/gin-gonic/gin"
)

func GetPublishList(c *gin.Context) {

	var queryVar struct {
		Token  string `json:"token" form:"token"`
		UserId int64  `json:"user_id" form:"user_id"`
	}

	//TODO错误处理
	if err := c.BindQuery(&queryVar); err != nil {
		SendResponseV(c, errno.ConvertErr(err), nil)
	}

	req := &video.GetPublishListRequest{Token: queryVar.Token, UserId: queryVar.UserId}

	videoList, err := rpc.GetPublishList(context.Background(), req)

	if err != nil {
		SendResponseV(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponseV(c, errno.Success, videoList)
}
