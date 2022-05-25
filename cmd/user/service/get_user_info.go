package service

import (
	"context"
	"github.com/douyin/cmd/user/dal/db"
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

// GetUserInfo query db to find information of a user
func (s *UserLoginService) GetUserInfo(req *user.DouyinUserRequest) (*db.UserInfo, error) {
	userInfo, err := db.GetUserInfo(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
