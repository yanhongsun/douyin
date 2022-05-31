package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/relation"
)

type GetFollowerListService struct {
	ctx context.Context
}

func NewGetFollowerListService(ctx context.Context) *GetFollowerListService {
	return &GetFollowerListService{ctx: ctx}
}

func (s *GetFollowerListService) GetFollowerList(req *relation.GetFollowerListRequest) ([]db.UserList, error) {

	return db.GetFollows(s.ctx, req.UserId)
}
