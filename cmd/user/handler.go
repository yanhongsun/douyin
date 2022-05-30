package main

import (
	"context"
	"douyin/cmd/user/global"
	"douyin/cmd/user/pack"
	"douyin/cmd/user/service"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"fmt"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuildRegisterResp(errno.ParamErr, -1, "")
		return resp, nil
	}

	userID, err := service.NewUserRegisterService(ctx).CreateUser(req)
	if err != nil {
		resp = pack.BuildRegisterResp(err, -1, "")
		return resp, nil
	}

	token, _ := global.Jwt.CreateToken(global.CustomClaims{
		Id: userID,
	})

	resp = pack.BuildRegisterResp(nil, userID, token)
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuildLoginResp(errno.ParamErr, -1, "")
		return resp, nil
	}

	userID, err := service.NewUserLoginService(ctx).CheckUser(req)
	if err != nil {
		resp = pack.BuildLoginResp(err, -1, "")
		return resp, nil
	}

	token, _ := global.Jwt.CreateToken(global.CustomClaims{
		Id: userID,
	})

	resp = pack.BuildLoginResp(nil, userID, token)
	return resp, nil
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	claim, err := global.Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildGetUserResp(err, nil)
		return resp, nil
	} else if claim.Id != int64(req.UserId) {
		resp = pack.BuildGetUserResp(errno.TokenInvalidErr, nil)
		return resp, nil
	}

	userInfo, err := service.NewGetUserInfoService(ctx).GetUserInfo(req)
	fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&=====", userInfo)
	if err != nil {
		resp = pack.BuildGetUserResp(err, nil)
		return resp, nil
	}
	resp = pack.BuildGetUserResp(nil, userInfo)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>resp", resp)
	return resp, nil
}
