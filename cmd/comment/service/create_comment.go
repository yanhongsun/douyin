package service

import (
	"context"
	"douyin/cmd/comment/dal/db"
	"douyin/cmd/comment/pack"
	"douyin/kitex_gen/comment"
	"time"
)

type CreateCommentService struct {
	ctx context.Context
}

func NewCreateCommentService(ctx context.Context) *CreateCommentService {
	return &CreateCommentService{ctx: ctx}
}

func (s *CreateCommentService) CreateComment(req *comment.CreateCommentRequest) (*comment.Comment, error) {
	snowflake := pack.Snowflake{
		Timestamp:    time.Now().UnixNano() / 1000000,
		Sequence:     0,
		Datacenterid: 12,
		Workerid:     2,
	}

	commentId, err := snowflake.NextVal()

	if err != nil {
		return nil, err
	}

	commentModel := db.Comment{
		CommentID: commentId,
		VedioID:   req.VedioId,
		UserID:    req.UserId,
		State:     true,
		Content:   req.Content,
	}

	return &comment.Comment{
		CommentId:  commentId,
		UserId:     req.UserId,
		Content:    req.Content,
		CreateDate: commentModel.Model.CreatedAt.String(),
	}, db.CreateComment(s.ctx, commentModel)
}
