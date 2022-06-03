package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/relation"
	"douyin/pkg/errno"

	"douyin/middleware"

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
		SendResponseRelation(c, errno.ConvertErr(err), nil)
		return
	}

	if queryVar.Token != "" {
		_, claims, err := middleware.ParseToken(queryVar.Token)
		if err != nil {
			SendResponseRelation(c, errno.ConvertErr(err), nil)
			return
		}
		queryVar.UserId = claims.UserID
	} else {
		SendResponseRelation(c, errno.ParamErr, nil)
		return
	}
	// fmt.Println(queryVar)
	//TODO错误处理
	if queryVar.UserId < 0 {
		SendResponseRelation(c, errno.ParamErr, nil)
		return
	}

	req := relation.RelationActionRequest{
		UserId:     queryVar.UserId,
		Token:      queryVar.Token,
		ToUserId:   queryVar.ToUserId,
		ActionType: queryVar.ActionType,
	}

	reqIsFollow := relation.IsFollowRequest{
		UserId:   queryVar.UserId,
		Token:    queryVar.Token,
		ToUserId: queryVar.ToUserId,
	}
	tag, err := rpc.IsFollow(context.Background(), &reqIsFollow)
	// fmt.Println("relation_action.go   关注了吗:")
	// fmt.Println(tag)
	if err != nil {
		SendBaseResponse(c, errno.ConvertErr(err))
		return
	}
	if queryVar.ActionType == 1 && tag {
		//fmt.Println("queryVar.ActionType==1  true")
		SendBaseResponse(c, errno.Success)
		return
	}
	if queryVar.ActionType == 2 && !tag {
		//fmt.Println("queryVar.ActionType==2  false")
		SendBaseResponse(c, errno.Success)
		return
	}
	err = rpc.RelationAction(context.Background(), &req)
	if err != nil {
		SendBaseResponse(c, errno.ConvertErr(err))
		return
	}
	SendBaseResponse(c, errno.Success)
}
