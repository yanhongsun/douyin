package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/like"
	"douyin/pkg/errno"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Likeyou like you action
func Likeyou(c *gin.Context) {
	var likeVar LikeyouParam
	if err := c.ShouldBind(&likeVar); err != nil {
		SendResponseThumb(c, errno.ConvertErr(err), nil)
		return
	}

	//如果为空则，上一步就return了？
	if likeVar.VideoId == 0 || likeVar.ActionType == 0 {
		SendResponseThumb(c, errno.ParamErr, nil)
		return
	}

	userID := likeVar.UserId

	err := rpc.Likeyou(context.Background(), &like.LikeyouRequest{
		UserId:     userID,
		Token:      "",
		VideoId:    likeVar.VideoId,
		ActionType: likeVar.ActionType,
	})
	if err != nil {
		SendResponseThumb(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponseThumb(c, errno.Success, nil)
}

func ThumbList(c *gin.Context) {
	var queryVar struct {
		UserId int64  `json:"user_id" form:"user_id"`
		Token  string `json:"token" form:"token"`
	}

	if err := c.BindQuery(&queryVar); err != nil {
		SendResponseThumb(c, errno.ConvertErr(err), nil)
		return
	}

	req := &like.ThumbListRequest{
		UserId: queryVar.UserId,
		Token:  queryVar.Token,
	}
	fmt.Println("即将开始调用rpc.ThumbList")
	videos, err := rpc.ThumbList(context.Background(), req)
	if err != nil {
		SendResponseThumb(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponseThumb(c, errno.Success, map[string]interface{}{"video_list": videos})
}
