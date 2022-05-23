package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/douyin/cmd/user/dal/db"
	"github.com/douyin/kitex_gen/douyin_user"
	"io"
)

type UserLoginService struct {
	ctx context.Context
}

func NewUserLoginService(ctx context.Context) *UserLoginService {
	return &UserLoginService{
		ctx: ctx,
	}
}

func (s *UserLoginService) UserLogin(req *douyin_user.DouyinUserLoginRequest) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprint("%x", h.Sum(nil))
	userName := req.Username
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, nil
	}
	u := users[0]
	if u.Password != passWord {
		return 0, nil
	}
	return int64(u.ID), nil
}
