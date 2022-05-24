package service

import (
	"context"
	"github.com/douyin/kitex_gen/user"
)

type UserLoginService struct {
	ctx context.Context
}

func NewUserLoginService(ctx context.Context) *UserLoginService {
	return &UserLoginService{
		ctx: ctx,
	}
}

func (s *UserLoginService) CheckUser(req *user.DouyinUserLoginRequest) error {
	// 返回id token
	return nil
}
