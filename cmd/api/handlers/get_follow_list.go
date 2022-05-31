package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/relation"
	"douyin/pkg/errno"

	"github.com/gin-gonic/gin"
)

func GetFollowList(c *gin.Context) {
	var queryVar struct {
		UserId int64  `json:"user_id" form:"user_id"`
		Token  string `json:"token" form:"token"`
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
	/*
		type GetFollowListRequest struct {
			UserId int64  `thrift:"user_id,1,required" json:"user_id"`
			Token  string `thrift:"token,2,required" json:"token"`
		}
	*/

	req := relation.GetFollowListRequest{
		UserId: queryVar.UserId,
		Token:  queryVar.Token,
	}
	user_list, err := rpc.GetFollowList(context.Background(), &req)
	if err != nil {
		SendResponseRelation(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponseRelation(c, errno.Success, user_list)
}
