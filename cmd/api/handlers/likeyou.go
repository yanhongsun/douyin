package handlers

import (
	"context"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/yanhongsun/douyin/cmd/api/rpc"
	"github.com/yanhongsun/douyin/kitex_gen/like"
	"github.com/yanhongsun/douyin/pkg/constants"
	"github.com/yanhongsun/douyin/pkg/errno"
)

// Likeyou like you action
func Likeyou(c *gin.Context) {
	var likeVar LikeyouParam
	if err := c.ShouldBind(&likeVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	//如果为空则，上一步就return了？
	if likeVar.VideoId == 0 || likeVar.ActionType == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))
	//todo: Token

	err := rpc.Likeyou(context.Background(), &like.LikeyouRequest{
		UserId:     userID,
		Token:      "",
		VideoId:    likeVar.VideoId,
		ActionType: likeVar.ActionType,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

func ThumbList(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))

	fmt.Println("claims:", claims)

	req := &like.ThumbListRequest{
		UserId: userID,
		Token:  "",
	}
	videos, err := rpc.ThumbList(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, map[string]interface{}{"video_list": videos})
}
