package main

import (
	"context"
	"douyin/cmd/comment/service"
	"douyin/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CreateComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CreateComment(ctx context.Context, req *comment.CreateCommentRequest) (resp *comment.CreateCommentResponse, err error) {
	resp = new(comment.CreateCommentResponse)

	if req.UserId <= 0 || req.VedioId <= 0 || len(req.Content) == 0 {
		resp.BaseResp.StatusCode = 10002
		resp.BaseResp.StatusMessage = "Wrong Parameter has been given"
		return resp, nil
	}

	comment, err := service.NewCreateCommentService(ctx).CreateComment(req)
	if err != nil {
		resp.BaseResp.StatusCode = 10001
		resp.BaseResp.StatusMessage = err.Error()
	}

	resp.BaseResp.StatusCode = 0
	resp.BaseResp.StatusMessage = "Success"
	resp.Comment = comment
	return resp, nil
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

// CreateCommentIndex implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CreateCommentIndex(ctx context.Context, req *comment.CreateCommentIndexRequset) (resp *comment.CreateCommentIndexResponse, err error) {
	// TODO: Your code here...
	return
}
