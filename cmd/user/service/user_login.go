package service

import (
	"context"
	"github.com/douyin/cmd/user/dal/db"
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

// CheckUser call db to check user is valid or not
func (s *UserLoginService) CheckUser(req *user.DouyinUserLoginRequest) (int64, string, error) {
	// TODO: 返回id token
	userToken, err := db.CheckUser(s.ctx, req.Username, req.Password)
	if err != nil {
		return -1, "", err
	}
	return userToken.UserID, userToken.Token, nil
}
