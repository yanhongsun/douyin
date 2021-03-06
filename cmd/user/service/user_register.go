package service

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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
	users, err := db.QueryUserByName(s.ctx, req.Username)
	if err != nil {
		return -1, err
	}
	if len(users) != 0 {
		return -1, errno.UserAlreadyExistErr
	}

	pwByte := []byte(req.Password)
	pwByte, _ = bcrypt.GenerateFromPassword(pwByte, bcrypt.DefaultCost)
	password := string(pwByte)
	fmt.Println(password, " ", len(password))

	res, err := db.CreateUser(s.ctx, req.Username, password)
	if err != nil {
		return -1, err
	}
	u := res[0]

	return u.ID, nil
}
