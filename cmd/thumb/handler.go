package main

import (
	"context"
	"github.com/yanhongsun/douyin/kitex_gen/like"
)

// ThumbServiceImpl implements the last service interface defined in the IDL.
type ThumbServiceImpl struct{}

// Likeyou implements the ThumbServiceImpl interface.
func (s *ThumbServiceImpl) Likeyou(ctx context.Context, request *like.LikeyouRequest) (resp *like.LikeyouResponse, err error) {
	// TODO: Your code here...

	return
}

// ThumbList implements the ThumbServiceImpl interface.
func (s *ThumbServiceImpl) ThumbList(ctx context.Context, request *like.ThumbListResponse) (resp *like.ThumbListResponse, err error) {
	// TODO: Your code here...
	return
}
