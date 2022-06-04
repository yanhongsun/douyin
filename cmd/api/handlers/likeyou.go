package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/like"
	"douyin/middleware"
	"douyin/pkg/errno"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Likeyou like you action
func Likeyou(c *gin.Context) {
	var likeVar LikeyouParam

	likeVar.ActionType = c.Query("action_type")
	likeVar.Token = c.Query("token")
	likeVar.VideoId = c.Query("video_id")

	fmt.Fprintln(gin.DefaultWriter, likeVar)

	videoId, err1 := strconv.ParseInt(likeVar.VideoId, 10, 64)
	actionType, err2 := strconv.ParseInt(likeVar.ActionType, 10, 64)

	//如果为空则，上一步就return了？
	if videoId == 0 || actionType == 0 || err1 != nil || err2 != nil {
		SendResponseThumb(c, errno.ParamErr, nil)
		return
	}

	_, claims, err := middleware.ParseToken(likeVar.Token)
	if err != nil {
		SendResponseThumb(c, errno.ConvertErr(err), nil)
		return
	}
	userID := claims.UserID

	err = rpc.Likeyou(context.Background(), &like.LikeyouRequest{
		UserId:     userID,
		Token:      "",
		VideoId:    videoId,
		ActionType: actionType,
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
