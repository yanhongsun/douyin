package service

import (
	"context"
	"douyin/cmd/relation/dal/redisCache"
	"douyin/kitex_gen/relation"
)

type IsFollowService struct {
	ctx context.Context
}

func NewIsFollowService(ctx context.Context) *IsFollowService {
	return &IsFollowService{ctx: ctx}
}

func (s *IsFollowService) IsFollow(req *relation.IsFollowRequest) (bool, error) {
	//fmt.Println("进入db，开始判断是否关注！")
	return redisCache.IsFollowed(s.ctx, req.UserId, req.ToUserId)
}
