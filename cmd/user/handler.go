package main

import (
	"context"
	"douyin/cmd/user/pack"
	"douyin/cmd/user/service"
	"douyin/kitex_gen/user"
	"douyin/middleware"
	"douyin/pkg/errno"
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

	// token, _ := global.Jwt.CreateToken(userID, global.JWTSetting.AppKey, global.JWTSetting.AppSecret)
	token, err := middleware.CreateToken(userID)
	if err != nil {
		resp = pack.BuildRegisterResp(err, -1, "")
		return resp, nil
	}

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

	// token, _ := global.Jwt.CreateToken(userID, global.JWTSetting.AppKey, global.JWTSetting.AppSecret)
	token, err := middleware.CreateToken(userID)
	if err != nil {
		resp = pack.BuildLoginResp(err, -1, "")
		return resp, nil
	}

	resp = pack.BuildLoginResp(nil, userID, token)
	return resp, nil
}

// IsUserExisted implements the UserServiceImpl interface.
func (s *UserServiceImpl) IsUserExisted(ctx context.Context, req *user.DouyinUserExistRequest) (resp *user.DouyinUserExistResponse, err error) {
	if req.TargetId == 0 {
		resp = pack.BuildUserExistResp(errno.ParamErr, false)
		return resp, nil
	}
	res, err := service.NewUserExistService(ctx).UserExist(req)
	if err != nil {
		resp = pack.BuildUserExistResp(err, false)
		return resp, nil
	}
	resp = pack.BuildUserExistResp(err, res)
	return resp, nil
}

// QueryCurUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryCurUser(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	if req.UserId == 0 {
		resp = pack.BuildQueryUserResp(errno.ParamErr, nil)
	}
	res, err := service.NewQueryUserService(ctx).QueryCurUserByID(req)
	if err != nil {
		resp = pack.BuildQueryUserResp(err, nil)
		return resp, nil
	}
	resp = pack.BuildQueryUserResp(nil, res)

	return resp, nil
}

// QueryOtherUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryOtherUser(ctx context.Context, req *user.DouyinQueryUserRequest) (resp *user.DouyinUserResponse, err error) {
	if req.UserId == 0 || req.TargetId == 0 {
		resp = pack.BuildQueryUserResp(errno.ParamErr, nil)
	}
	res, err := service.NewQueryUserService(ctx).QueryOtherUserByID(req)
	if err != nil {
		resp = pack.BuildQueryUserResp(err, nil)
		return resp, nil
	}
	resp = pack.BuildQueryUserResp(nil, res)

	return resp, nil
}

// MultiQueryUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MultiQueryUser(ctx context.Context, req *user.DouyinMqueryUserRequest) (resp *user.DouyinMqueryUserResponse, err error) {
	res, err := service.NewMultiQueryUserService(ctx).MultiQueryUser(req)
	if err != nil {
		resp = pack.BuildMultiQueryUserResp(err, nil)
		return resp, nil
	}

	resp = pack.BuildMultiQueryUserResp(nil, res)

	return resp, nil
}
