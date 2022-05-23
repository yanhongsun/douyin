package service

import (
	"context"
	"github.com/douyin/cmd/user/dal/db"
	"github.com/douyin/kitex_gen/douyin_user"
)

type UserRegisterService struct {
	ctx context.Context
}

func NewUserRegisterService(ctx context.Context) *UserRegisterService {
	return &UserRegisterService{
		ctx: ctx,
	}
}

func (s *UserRegisterService) UserRegister(req *douyin_user.DouyinUserRegisterRequest) error {
	users, err := db.QueryUser
}
