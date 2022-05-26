package service

import (
	"context"
	"douyin/cmd/comment/dal/db"
	"douyin/cmd/comment/pack"
	"douyin/kitex_gen/comment"
)

type QueryCommentsService struct {
	ctx context.Context
}

func NewQueryCommentsService(ctx context.Context) *QueryCommentsService {
	return &QueryCommentsService{ctx: ctx}
}

func (s *QueryCommentsService) QueryComments(req *comment.QueryCommentsRequest) ([]*comment.Comment, error) {
	res, err := db.QueryComments(s.ctx, req.VideoId, 10000, 0)

	if err != nil {
		return nil, err
	}

	return pack.ChangeComments(res), nil
}
