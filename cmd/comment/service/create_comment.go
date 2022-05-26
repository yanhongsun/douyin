package service

import (
	"context"
	"douyin/cmd/comment/dal/db"
	"douyin/kitex_gen/comment"

	"douyin/cmd/comment/pack/snowflake"
)

var snowflakeNode *snowflake.Node

func InitSnowflakeNode() {
	tmpNode, err := snowflake.NewNode(1)
	if err != nil {
		panic("snowflake error")
	}
	snowflakeNode = tmpNode
}

type CreateCommentService struct {
	ctx context.Context
}

func NewCreateCommentService(ctx context.Context) *CreateCommentService {
	return &CreateCommentService{ctx: ctx}
}

func (s *CreateCommentService) CreateComment(req *comment.CreateCommentRequest) (*comment.Comment, error) {

	commentId := snowflakeNode.Generate().Int64()

	commentModel := db.Comment{
		CommentID: commentId,
		VideoID:   req.VideoId,
		UserID:    req.UserId,
		State:     true,
		Content:   req.Content,
	}

	err := db.CreateComment(s.ctx, &commentModel)

	return &comment.Comment{
		CommentId:  commentId,
		UserId:     req.UserId,
		Content:    req.Content,
		CreateDate: commentModel.Model.CreatedAt.Format("01-02"),
	}, err
}
