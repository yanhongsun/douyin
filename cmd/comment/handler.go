package main

import (
	"context"
	"douyin/cmd/comment/pack"
	"douyin/cmd/comment/rpc"
	"douyin/cmd/comment/service"
	"douyin/kitex_gen/comment"
	"douyin/pkg/errno"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CreateComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CreateComment(ctx context.Context, req *comment.CreateCommentRequest) (resp *comment.CreateCommentResponse, err error) {
	resp = new(comment.CreateCommentResponse)

	exist, err := rpc.VerifyVideoId(ctx, req.VideoId, "123")
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.CommentServiceErr.WithMessage(err.Error()))
		return resp, nil
	}
	if !exist {
		resp.BaseResp = pack.BuildBaseResp(errno.VideoIdErr)
		return resp, nil
	}

	if len(req.Content) == 0 && len(req.Content) <= 20000 {
		resp.BaseResp = pack.BuildBaseResp(errno.CommentParamErr)
		return resp, nil
	}

	createcomment, err := service.NewCreateCommentService(ctx).CreateComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Comment = createcomment

	return resp, nil
}

// DeleteComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) DeleteComment(ctx context.Context, req *comment.DeleteCommentRequest) (resp *comment.DeleteCommentResponse, err error) {
	resp = new(comment.DeleteCommentResponse)

	exist, err := rpc.VerifyVideoId(ctx, req.VideoId, "123")
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.CommentServiceErr.WithMessage(err.Error()))
		return resp, nil
	}
	if !exist {
		resp.BaseResp = pack.BuildBaseResp(errno.VideoIdErr)
		return resp, nil
	}

	err = service.NewDeleteCommentService(ctx).DeleteComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)

	return resp, nil
}

// QueryComments implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) QueryComments(ctx context.Context, req *comment.QueryCommentsRequest) (resp *comment.QueryCommentsResponse, err error) {
	resp = new(comment.QueryCommentsResponse)

	exist, err := rpc.VerifyVideoId(ctx, req.VideoId, "123")
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.CommentServiceErr.WithMessage(err.Error()))
		return resp, nil
	}
	if !exist {
		resp.BaseResp = pack.BuildBaseResp(errno.VideoIdErr)
		return resp, nil
	}

	res, err := service.NewQueryCommentsService(ctx).QueryComments(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Comments = res

	return resp, nil
}

// QueryCommentNumber implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) QueryCommentNumber(ctx context.Context, req *comment.QueryCommentNumberRequest) (resp *comment.QueryCommentNumberResponse, err error) {
	resp = new(comment.QueryCommentNumberResponse)

	res, err := service.NewQueryCommentNumberService(ctx).QueryCommentNumber(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.CommentNumber = res

	return resp, nil
}
