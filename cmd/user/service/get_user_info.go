package service

import (
	"context"
	"github.com/douyin/cmd/user/dal/db"
	"github.com/douyin/cmd/user/pack"
	"github.com/douyin/kitex_gen/user"
)

type GetUserInfoService struct {
	ctx context.Context
}

func NewGetUserInfoService(ctx context.Context) *UserLoginService {
	return &UserLoginService{
		ctx: ctx,
	}
}

// QueryUser query db to find information of a user
func (s *UserLoginService) QueryUser(req *user.DouyinUserRequest) (*user.User, error) {
	userInfo, err := db.GetUserInfo(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.UserInfo(userInfo), nil
}
