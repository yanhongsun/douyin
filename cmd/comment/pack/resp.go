package pack

import (
	"douyin/kitex_gen/comment"
	"douyin/pkg/errno"
	"errors"
)

func BuildBaseResp(err error) *comment.BaseResp {
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

func baseResp(err errno.ErrNo) *comment.BaseResp {
	return &comment.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg}
}
