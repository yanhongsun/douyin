package service

import (
	"context"
	"douyin/cmd/comment/dal/db"
	"douyin/kitex_gen/comment"
)

type DeleteCommentService struct {
	ctx context.Context
}

func NewDeleteCommentService(ctx context.Context) *DeleteCommentService {
	return &DeleteCommentService{ctx: ctx}
}

func (s *DeleteCommentService) DeleteComment(req *comment.DeleteCommentRequest) error {
	return db.DeleteComment(s.ctx, req.CommentId, req.VedioId, req.UserId)
}
