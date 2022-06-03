package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/cmd/relation/dal/redisCache"
	"douyin/kitex_gen/relation"
)

type GetFollowerListService struct {
	ctx context.Context
}

func NewGetFollowerListService(ctx context.Context) *GetFollowerListService {
	return &GetFollowerListService{ctx: ctx}
}

func (s *GetFollowerListService) GetFollowerList(req *relation.GetFollowerListRequest) ([]db.UserList, error) {

	return redisCache.GetFans(s.ctx, req.UserId)
}
