package service

import (
	"context"
	"crypto/md5"
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"fmt"
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

// CheckUser call db to check user is valid or not
func (s *UserLoginService) CheckUser(req *user.DouyinUserLoginRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return -1, err
	}
	// salt, err := db.QuerySalt(s.ctx, req.Username)
	// passWord := fmt.Sprintf("%x", h.Sum(salt))
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.Username
	res, err := db.QueryUserByName(s.ctx, userName)
	if err != nil {
		return -1, err
	}
	if len(res) == 0 {
		return -1, errno.UserNotExistErr
	}
	u := res[0]
	if u.Password != passWord {
		return -1, errno.LoginErr
	}

	return int64(u.ID), nil
}
