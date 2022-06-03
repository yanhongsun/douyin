package service

import (
	"context"
	"douyin/cmd/thumb/dal/db"
	"douyin/cmd/thumb/pack"
	"douyin/kitex_gen/like"
	"fmt"
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
	fmt.Println("来到了service ThumbList中，req=", req)
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
