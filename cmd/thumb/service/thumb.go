package service

import (
	"context"
	"github.com/yanhongsun/douyin/kitex_gen/like"
)

type ThumbService struct {
	ctx context.Context
}
type IThumbService interface {
	ThumbList(req *like.ThumbListRequest) []*like.Video
	Likeyou(req *like.LikeyouRequest) error
}

// NewThumbService new ThumbService
func NewThumbService(ctx context.Context) *ThumbService {
	return &ThumbService{
		ctx: ctx,
	}
}

// 点赞操作

func (t ThumbService) ThumbList(req *like.ThumbListRequest) ([]*like.Video, error) {
	//TODO implement me
	panic("implement me")
}

func (t ThumbService) Likeyou(req *like.LikeyouRequest) error {
	//TODO implement me
	panic("implement me")
}
