package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/cmd/relation/dal/redisCache"
	"douyin/kitex_gen/relation"
)

type GetFollowListService struct {
	ctx context.Context
}

func NewGetFollowListService(ctx context.Context) *GetFollowListService {
	return &GetFollowListService{ctx: ctx}
}

func (s *GetFollowListService) GetFollowList(req *relation.GetFollowListRequest) ([]db.UserList, error) {

	return redisCache.GetFollows(s.ctx, req.UserId)
}
