package service

import (
	"context"
	"github.com/yanhongsun/douyin/cmd/thumb/dal/db"
	"github.com/yanhongsun/douyin/cmd/thumb/pack"
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
	videos, err := db.ListVideo(t.ctx, req.UserId)
	userInfo, err := db.GetUserInfo(t.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	pack.Videos(videos, userInfo)
	return pack.Videos(videos, userInfo), nil
}

//看笔记服务CreateNode怎么写的。。
func (t ThumbService) Likeyou(req *like.LikeyouRequest) error {
	err := db.UpdatdeVideo(t.ctx, req.UserId, req.VideoId, req.ActionType)
	if err != nil {
		return err
	}
	return nil
}
