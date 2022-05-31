package pack

import (
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"errors"
)

func BuildRegisterResp(err error, userID int64, token string) *user.DouyinUserRegisterResponse {
	var resp user.DouyinUserRegisterResponse
	if err == nil {
		resp.SetStatusCode(errno.Success.ErrCode)
		resp.SetStatusMsg(&errno.Success.ErrMsg)
		resp.SetUserId(userID)
		resp.SetToken(token)
		return &resp
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		resp.SetStatusCode(e.ErrCode)
		resp.SetStatusMsg(&e.ErrMsg)
		resp.SetUserId(userID)
		resp.SetToken(token)
		return &resp
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	resp.SetStatusCode(s.ErrCode)
	resp.SetStatusMsg(&s.ErrMsg)
	resp.SetUserId(userID)
	resp.SetToken(token)
	return &resp
}

func BuildLoginResp(err error, userID int64, token string) *user.DouyinUserLoginResponse {
	var resp user.DouyinUserLoginResponse
	if err == nil {
		resp.SetStatusCode(errno.Success.ErrCode)
		resp.SetStatusMsg(&errno.Success.ErrMsg)
		resp.SetUserId(userID)
		resp.SetToken(token)
		return &resp
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		resp.SetStatusCode(e.ErrCode)
		resp.SetStatusMsg(&e.ErrMsg)
		resp.SetUserId(userID)
		resp.SetToken(token)
		return &resp
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	resp.SetStatusCode(s.ErrCode)
	resp.SetStatusMsg(&s.ErrMsg)
	resp.SetUserId(userID)
	resp.SetToken(token)
	return &resp
}

func BuildUserExistResp(err error, isExist bool) *user.DouyinUserExistResponse {
	var resp user.DouyinUserExistResponse
	if err == nil {
		resp.SetStatusCode(errno.Success.ErrCode)
		resp.SetStatusMsg(&errno.Success.ErrMsg)
		resp.SetIsExisted(isExist)
		return &resp
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		resp.SetStatusCode(e.ErrCode)
		resp.SetStatusMsg(&e.ErrMsg)
		resp.SetIsExisted(isExist)
		return &resp
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	resp.SetStatusCode(s.ErrCode)
	resp.SetStatusMsg(&s.ErrMsg)
	resp.SetIsExisted(isExist)
	return &resp
}

func BuildQueryUserResp(err error, userInfo *user.User) *user.DouyinUserResponse {
	var resp user.DouyinUserResponse
	if err == nil {
		resp.SetStatusCode(errno.Success.ErrCode)
		resp.SetStatusMsg(&errno.Success.ErrMsg)
		resp.SetUser(userInfo)
		return &resp
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		resp.SetStatusCode(e.ErrCode)
		resp.SetStatusMsg(&e.ErrMsg)
		resp.SetUser(&user.User{
			Id: -1, Name: "", FollowCount: nil, FollowerCount: nil, IsFollow: false,
		})
		return &resp
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	resp.SetStatusCode(s.ErrCode)
	resp.SetStatusMsg(&s.ErrMsg)
	resp.SetUser(&user.User{
		Id: -1, Name: "", FollowCount: nil, FollowerCount: nil, IsFollow: false,
	})
	return &resp
}
