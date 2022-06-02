package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/relation"
	"fmt"
)

type FollowService struct {
	ctx context.Context
}

func NewFollowService(ctx context.Context) *FollowService {
	return &FollowService{ctx: ctx}
}

func (s *FollowService) Follow(req *relation.RelationActionRequest) error {
	fmt.Println("进入数据库，开始关注")
	return db.Follow(s.ctx, req.UserId, req.ToUserId)
}
