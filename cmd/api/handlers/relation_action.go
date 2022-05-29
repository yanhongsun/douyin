package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/relation"
	"douyin/pkg/errno"

	"github.com/gin-gonic/gin"
)

func RelationAction(c *gin.Context) {
	var queryVar struct {
		UserId     int64  `json:"user_id" form:"user_id"`
		Token      string `json:"token" form:"token"`
		ToUserId   int64  `json:"to_user_id" form:"to_user_id"`
		ActionType int32  `json:"action_type" form:"action_type"`
	}
	//TODO错误处理
	if err := c.BindQuery(&queryVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	//TODO错误处理
	if queryVar.UserId < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	/*
		type RelationActionRequest struct {
		UserId     int64  `thrift:"user_id,1,required" json:"user_id"`
		Token      string `thrift:"token,2,required" json:"token"`
		ToUserId   int64  `thrift:"to_user_id,3,required" json:"to_user_id"`
		ActionType int32  `thrift:"action_type,4,required" json:"action_type"`
	}*/
	req := relation.RelationActionRequest{
		UserId:     queryVar.UserId,
		Token:      queryVar.Token,
		ToUserId:   queryVar.ToUserId,
		ActionType: queryVar.ActionType,
	}

	err := rpc.RelationAction(context.Background(), &req)
	if err != nil {
		SendBaseResponse(c, errno.ConvertErr(err))
		return
	}
	SendBaseResponse(c, errno.Success)
}
