package service

import (
	"context"
	"douyin/cmd/relation/dal/redisCache"
	"douyin/kitex_gen/relation"
)

type UnFollowService struct {
	ctx context.Context
}

func NewUnFollowService(ctx context.Context) *UnFollowService {
	return &UnFollowService{ctx: ctx}
}

func (s *UnFollowService) UnFollow(req *relation.RelationActionRequest) error {

	return redisCache.UnFollow(s.ctx, req.UserId, req.ToUserId)
}
