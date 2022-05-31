package service

import (
	"context"
	"douyin/cmd/comment/repository"
	"douyin/kitex_gen/comment"
	"strconv"

	"golang.org/x/sync/singleflight"
)

var gDeleteComment singleflight.Group

type DeleteCommentService struct {
	ctx context.Context
}

func NewDeleteCommentService(ctx context.Context) *DeleteCommentService {
	return &DeleteCommentService{ctx: ctx}
}

func (s *DeleteCommentService) DeleteComment(req *comment.DeleteCommentRequest) error {
	key := strconv.FormatInt(req.CommentId, 10)

	_, err, _ := gDeleteComment.Do(key, func() (interface{}, error) {
		return nil, repository.ProducerComment(s.ctx, 2, nil, req.CommentId, req.VideoId, nil, req.UserId)
	})

	return err
}
