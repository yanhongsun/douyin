package rpc

import (
	"context"
	"github.com/douyin/cmd/api/handlers"
	"github.com/douyin/kitex_gen/user"
	"github.com/douyin/kitex_gen/user/userservice"
	"github.com/douyin/pkg/errno"
)

var userClient userservice.Client

// CreateUser 创建新用户
func CreateUser(ctx context.Context, req *user.DouyinUserRegisterRequest) (int64, string, error) {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return -1, "", err
	}

	if resp.StatusCode != 0 {
		return -1, "", errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}

	return resp.UserId, resp.Token, nil
}

// CheckUser 检查用户是否存在
func CheckUser(ctx context.Context, req *user.DouyinUserLoginRequest) (int64, string, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, "", err
	}

	if resp.StatusCode != 0 {
		return 0, "", errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}

	return resp.UserId, resp.Token, nil
}

// GetUserInfo 获取用户信息
func GetUserInfo(ctx context.Context, req *user.DouyinUserRequest) (*handlers.UserInfo, error) {
	resp, err := userClient.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	// 正常, 返回用户信息
	var userInfo handlers.UserInfo
	userInfo.UserID = resp.User.Id
	userInfo.Username = resp.User.Name
	userInfo.FollowCount = resp.User.GetFollowCount()
	userInfo.FollowerCount = resp.User.GetFollowerCount()
	userInfo.IsFollow = resp.User.IsFollow
	// TODO: type ok?
	return &userInfo, nil
}
