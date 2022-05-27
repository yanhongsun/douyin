package service

import (
	"context"
	"douyin/cmd/comment/dal/mysqldb"
	"douyin/kitex_gen/comment"
)

type QueryCommentNumberService struct {
	ctx context.Context
}

func NewQueryCommentNumberService(ctx context.Context) *QueryCommentNumberService {
	return &QueryCommentNumberService{ctx: ctx}
}

func (s *QueryCommentNumberService) QueryCommentNumber(req *comment.QueryCommentNumberRequest) (int64, error) {
	res, err := mysqldb.QueryCommentsNumber(s.ctx, req.VideoId)

	if err != nil {
		return 0, err
	}

	return res, nil
}
