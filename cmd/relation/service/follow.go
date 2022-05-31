package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/relation"
)

type FollowService struct {
	ctx context.Context
}

func NewFollowService(ctx context.Context) *FollowService {
	return &FollowService{ctx: ctx}
}

func (s *FollowService) Follow(req *relation.RelationActionRequest) error {

	return db.Follow(s.ctx, req.UserId, req.ToUserId)
}
