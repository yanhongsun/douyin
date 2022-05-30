package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/relation"
)

type IsFollowService struct {
	ctx context.Context
}

func NewIsFollowService(ctx context.Context) *IsFollowService {
	return &IsFollowService{ctx: ctx}
}

func (s *IsFollowService) IsFollow(req *relation.IsFollowRequest) (bool, error) {

	return db.IsFollowed(s.ctx, req.UserId, req.ToUserId)
}
