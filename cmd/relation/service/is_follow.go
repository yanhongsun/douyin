package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/relation"
	"fmt"
)

type IsFollowService struct {
	ctx context.Context
}

func NewIsFollowService(ctx context.Context) *IsFollowService {
	return &IsFollowService{ctx: ctx}
}

func (s *IsFollowService) IsFollow(req *relation.IsFollowRequest) (bool, error) {
	fmt.Println("进入db，开始判断是否关注！")
	return db.IsFollowed(s.ctx, req.UserId, req.ToUserId)
}
