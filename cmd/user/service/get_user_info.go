package service

import (
	"context"
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

func (s *UserLoginService) QueryUser(req *user.DouyinUserRequest) error {
	return nil
}
