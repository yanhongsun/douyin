package service

import (
	"context"
	"douyin/cmd/comment/dal/mysqldb"
	"douyin/cmd/comment/pack"
	"douyin/kitex_gen/comment"

	"douyin/pkg/snowflake"

	"douyin/cmd/comment/repository"
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

	commentModel := mysqldb.Comment{
		CommentID: commentId,
		VideoID:   req.VideoId,
		UserID:    req.UserId,
		State:     true,
		Content:   req.Content,
	}

	if err := repository.ProducerCreateComment(&commentModel); err != nil {
		return nil, err
	}

	return pack.ChangeComment(&commentModel), nil
}
