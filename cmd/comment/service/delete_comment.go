package service

import (
	"context"
	"douyin/cmd/comment/repository"
	"douyin/kitex_gen/comment"
)

type DeleteCommentService struct {
	ctx context.Context
}

func NewDeleteCommentService(ctx context.Context) *DeleteCommentService {
	return &DeleteCommentService{ctx: ctx}
}

func (s *DeleteCommentService) DeleteComment(req *comment.DeleteCommentRequest) error {
	return repository.ProducerDeleteComment(req.CommentId, req.VideoId, req.UserId)
}
