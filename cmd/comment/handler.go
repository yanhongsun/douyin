package main

import (
	"context"
	"douyin/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CreateComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CreateComment(ctx context.Context, req *comment.CreateCommentRequest) (resp *comment.CreateCommentResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) DeleteComment(ctx context.Context, req *comment.DeleteCommentRequest) (resp *comment.DeleteCommentResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryComments implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) QueryComments(ctx context.Context, req *comment.QueryCommentsRequest) (resp *comment.QueryCommentsResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryCommentNumber implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) QueryCommentNumber(ctx context.Context, req *comment.QueryCommentNumberRequest) (resp *comment.QueryCommentNumberResponse, err error) {
	// TODO: Your code here...
	return
}
