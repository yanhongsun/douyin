package service

import (
	"context"
	"github.com/douyin/cmd/user/dal/db"
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

// CreateUser call db to create a user
func (s *UserRegisterService) CreateUser(req *user.DouyinUserRegisterRequest) (int64, string, error) {
	userToken, err := db.CreateUser(s.ctx, req.Username, req.Password)
	if err != nil {
		return -1, "", err
	}
	return userToken.UserID, userToken.Token, nil
}
