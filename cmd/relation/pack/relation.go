package pack

import (
	"douyin/kitex_gen/relation"
	"douyin/pkg/errno"
	"errors"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *relation.BaseResponse {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *relation.BaseResponse {
	var resp relation.BaseResponse
	resp.SetStatusCode(err.ErrCode)
	resp.SetStatusMsg(&err.ErrMsg)
	return &(resp)
}
