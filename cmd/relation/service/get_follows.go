package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/relation"
	"fmt"
)

type GetFollowListService struct {
	ctx context.Context
}

func NewGetFollowListService(ctx context.Context) *GetFollowListService {
	return &GetFollowListService{ctx: ctx}
}

func (s *GetFollowListService) GetFollowList(req *relation.GetFollowListRequest) ([]db.UserList, error) {
	fmt.Println("进入db")
	return db.GetFollows(s.ctx, req.UserId)
}
