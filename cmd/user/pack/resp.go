package pack

import (
	"errors"
	"github.com/douyin/kitex_gen/user"
	"github.com/douyin/pkg/errno"
)

func BuildRegisterResp(err error, userID int64, token string) *user.DouyinUserRegisterResponse {
	var resp *user.DouyinUserRegisterResponse
	if err == nil {
		resp.SetStatusCode(errno.Success.ErrCode)
		resp.SetStatusMsg(&errno.Success.ErrMsg)
		resp.SetUserId(userID)
		resp.SetToken(token)
		return resp
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		resp.SetStatusCode(e.ErrCode)
		resp.SetStatusMsg(&e.ErrMsg)
		resp.SetUserId(userID)
		resp.SetToken(token)
		return resp
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	resp.SetStatusCode(s.ErrCode)
	resp.SetStatusMsg(&s.ErrMsg)
	resp.SetUserId(userID)
	resp.SetToken(token)
	return resp
}

func BuildLoginResp(err error, userID int64, token string) *user.DouyinUserLoginResponse {
	var resp *user.DouyinUserLoginResponse
	if err == nil {
		resp.SetStatusCode(errno.Success.ErrCode)
		resp.SetStatusMsg(&errno.Success.ErrMsg)
		resp.SetUserId(userID)
		resp.SetToken(token)
		return resp
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		resp.SetStatusCode(e.ErrCode)
		resp.SetStatusMsg(&e.ErrMsg)
		resp.SetUserId(userID)
		resp.SetToken(token)
		return resp
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	resp.SetStatusCode(s.ErrCode)
	resp.SetStatusMsg(&s.ErrMsg)
	resp.SetUserId(userID)
	resp.SetToken(token)
	return resp
}

func BuildGetUserResp(err error, userInfo *user.User) *user.DouyinUserResponse {
	var resp *user.DouyinUserResponse
	if err == nil {
		resp.SetStatusCode(errno.Success.ErrCode)
		resp.SetStatusMsg(&errno.Success.ErrMsg)
		resp.SetUser(nil)
		return resp
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		resp.SetStatusCode(e.ErrCode)
		resp.SetStatusMsg(&e.ErrMsg)
		resp.SetUser(nil)
		return resp
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	resp.SetStatusCode(s.ErrCode)
	resp.SetStatusMsg(&s.ErrMsg)
	resp.SetUser(userInfo)
	return resp
}
