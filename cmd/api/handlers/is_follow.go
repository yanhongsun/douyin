package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/relation"
	"douyin/pkg/errno"

	"github.com/gin-gonic/gin"
)

func IsFollow(c *gin.Context) {
	var queryVar struct {
		UserId   int64  `json:"user_id" form:"user_id"`
		Token    string `json:"token" form:"token"`
		ToUserId int64  `json:"to_user_id" form:"to_user_id"`
	}
	//TODO错误处理
	if err := c.BindQuery(&queryVar); err != nil {
		SendResponseRelation(c, errno.ConvertErr(err), nil)
		return
	}
	//TODO错误处理
	if queryVar.UserId < 0 {
		SendResponseRelation(c, errno.ParamErr, nil)
		return
	}
	req := relation.IsFollowRequest{
		UserId:   queryVar.UserId,
		Token:    queryVar.Token,
		ToUserId: queryVar.ToUserId,
	}
	tag, err := rpc.IsFollow(context.Background(), &req)
	if err != nil {
		SendResponseRelation(c, errno.ConvertErr(err), false)
		return
	}
	SendResponseRelation(c, errno.Success, tag)
}
