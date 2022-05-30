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

type UserRegisterService struct {
	ctx context.Context
}

func NewUserRegisterService(ctx context.Context) *UserRegisterService {
	return &UserRegisterService{
		ctx: ctx,
	}
}

// CreateUser call db to create a user
func (s *UserRegisterService) CreateUser(req *user.DouyinUserRegisterRequest) (int64, error) {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return -1, err
	}
	if len(users) != 0 {
		return -1, errno.UserAlreadyExistErr
	}
	// crypt
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return -1, err
	}
	// TODO: set nil as salt?
	// var salt []byte
	// err = db.CreateSalt(s.ctx, req.Username, salt)
	// passWord := fmt.Sprintf("%x", h.Sum(salt))
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	res, err := db.CreateUser(s.ctx, req.Username, passWord)
	if err != nil {
		return -1, err
	}
	u := res[0]

	return int64(u.ID), nil
}
