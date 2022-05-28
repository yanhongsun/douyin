package handlers

import (
	"context"
	"fmt"
	"github.com/douyin/cmd/api/rpc"
	"github.com/douyin/kitex_gen/user"
	"github.com/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func QueryUser(c *gin.Context) {
	var queryVar UserInfoParam
	// fmt.Println(c.)
	if err := c.ShouldBind(&queryVar); err != nil {
		fmt.Println("........", queryVar)
		SendUserInfoResponse(c, errno.ConvertErr(err), &UserInfo{
			ID:            -1,
			Name:          "",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		})
		return
	}

	userInfo, err := rpc.GetUserInfo(context.Background(), &user.DouyinUserRequest{
		UserId: queryVar.UserID,
		Token:  queryVar.Token,
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
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++", userInfo)
	SendUserInfoResponse(c, errno.Success, &UserInfo{
		ID:            userInfo.ID,
		Name:          userInfo.Name,
		FollowCount:   userInfo.FollowCount,
		FollowerCount: userInfo.FollowerCount,
		IsFollow:      false,
	})
}
