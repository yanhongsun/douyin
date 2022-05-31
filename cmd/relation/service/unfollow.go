package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/relation"
)

type UnFollowService struct {
	ctx context.Context
}

func NewUnFollowService(ctx context.Context) *UnFollowService {
	return &UnFollowService{ctx: ctx}
}

func (s *UnFollowService) UnFollow(req *relation.RelationActionRequest) error {

	return db.Unfollow(s.ctx, req.UserId, req.ToUserId)
}
