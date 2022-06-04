package main

import (
	"context"
	"douyin/cmd/thumb/pack"
	"douyin/cmd/thumb/service"
	"douyin/kitex_gen/like"
	"douyin/pkg/errno"
	"fmt"
)

// ThumbServiceImpl implements the last service interface defined in the IDL.
type ThumbServiceImpl struct{}

// Likeyou implements the ThumbServiceImpl interface.
func (s *ThumbServiceImpl) Likeyou(ctx context.Context, request *like.LikeyouRequest) (resp *like.LikeyouResponse, err error) {
	if request.UserId == 0 || request.VideoId == 0 {
		return pack.BuildLikeyouResp(errno.ParamErr), nil
	}
	err = service.NewThumbService(ctx).Likeyou(request)
	// err是否为nil都兼容处理了
	return pack.BuildLikeyouResp(err), nil
}

// ThumbList implements the ThumbServiceImpl interface.
func (s *ThumbServiceImpl) ThumbList(ctx context.Context, request *like.ThumbListRequest) (resp *like.ThumbListResponse, err error) {
	if request.UserId == 0 {
		return pack.BuildThumblistResp(nil, errno.ParamErr), err
	}
	list, err := service.NewThumbService(ctx).ThumbList(request)
	return pack.BuildThumblistResp(list, err), nil
}
