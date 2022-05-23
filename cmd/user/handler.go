package main

import (
	"context"
	"github.com/douyin/kitex_gen/douyin_user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *douyin_user.DouyinUserRegisterRequest) (resp *douyin_user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *douyin_user.DouyinUserLoginRequest) (resp *douyin_user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *douyin_user.DouyinUserRequest) (resp *douyin_user.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}
