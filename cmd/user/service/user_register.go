package service

import (
	"context"
	"github.com/douyin/kitex_gen/user"
)

type UserRegisterService struct {
	ctx context.Context
}

func NewUserRegisterService(ctx context.Context) *UserRegisterService {
	return &UserRegisterService{
		ctx: ctx,
	}
}

func (s *UserRegisterService) CreateUser(req *user.DouyinUserRegisterRequest) error {
	return nil
}
