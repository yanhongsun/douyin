package handlers

// TODO: remove this file later

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

type IsUserExistedParam struct {
	UserID int64  `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
}

func IsUserExisted(c *gin.Context) {
	var existVar IsUserExistedParam

	if err := c.BindQuery(&existVar); err != nil {
		SendIsUserExistedResponse(c, errno.ConvertErr(err), false)
		return
	}

	isExisted, err := rpc.IsUserExisted(context.Background(), &user.DouyinUserExistRequest{
		TargetId: existVar.UserID,
	})

	if err != nil {
		SendIsUserExistedResponse(c, errno.ConvertErr(err), false)
		return
	}

	SendIsUserExistedResponse(c, errno.Success, isExisted)
}

func QueryOtherUser(c *gin.Context) {
	var queryVar struct {
		UserID   int64  `json:"user_id" form:"user_id"`
		TargetID int64  `json:"target_id" form:"target_id"`
		Token    string `json:"token" form:"token"`
	}
	if err := c.BindQuery(&queryVar); err != nil {
		SendUserInfoResponse(c, errno.ConvertErr(err), &UserInfo{
			ID:            -1,
			Name:          "",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		})
		return
	}
	userInfo, err := rpc.QueryOtherUser(context.Background(), &user.DouyinQueryUserRequest{
		UserId:   queryVar.UserID,
		TargetId: queryVar.TargetID,
		Token:    queryVar.Token,
	})
	if err != nil {
		SendUserInfoResponse(c, errno.ConvertErr(err), &UserInfo{
			ID:            -1,
			Name:          "",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		})
		return
	}
	SendUserInfoResponse(c, errno.Success, &UserInfo{
		ID:            userInfo.ID,
		Name:          userInfo.Name,
		FollowCount:   userInfo.FollowCount,
		FollowerCount: userInfo.FollowerCount,
		IsFollow:      false,
	})
}
