package pack

import (
	"errors"
	"github.com/douyin/kitex_gen/user"
	"github.com/douyin/pkg/errno"
)

func BuildRegisterResp(err error) *user.DouyinUserRegisterResponse {
	var resp *user.DouyinUserRegisterResponse
	if err == nil {
		resp.SetStatusCode(errno.Success.ErrCode)
		resp.SetStatusMsg(&errno.Success.ErrMsg)
		return resp
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		resp.SetStatusCode(e.ErrCode)
		resp.SetStatusMsg(&e.ErrMsg)
		return resp
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	resp.SetStatusCode(s.ErrCode)
	resp.SetStatusMsg(&s.ErrMsg)
	return resp
}

func BuildLoginResp(err error) *user.DouyinUserLoginResponse {
	var resp *user.DouyinUserLoginResponse
	if err == nil {
		resp.SetStatusCode(errno.Success.ErrCode)
		resp.SetStatusMsg(&errno.Success.ErrMsg)
		return resp
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		resp.SetStatusCode(e.ErrCode)
		resp.SetStatusMsg(&e.ErrMsg)
		return resp
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	resp.SetStatusCode(s.ErrCode)
	resp.SetStatusMsg(&s.ErrMsg)
	return resp
}

func BuildGetUserResp(err error) *user.DouyinUserResponse {
	var resp *user.DouyinUserResponse
	if err == nil {
		resp.SetStatusCode(errno.Success.ErrCode)
		resp.SetStatusMsg(&errno.Success.ErrMsg)
		return resp
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		resp.SetStatusCode(e.ErrCode)
		resp.SetStatusMsg(&e.ErrMsg)
		return resp
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	resp.SetStatusCode(s.ErrCode)
	resp.SetStatusMsg(&s.ErrMsg)
	return resp
}