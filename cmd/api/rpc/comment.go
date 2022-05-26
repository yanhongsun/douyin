package rpc

import (
	"context"
	"douyin/cmd/api/common"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/comment/commentservice"
	"douyin/pkg/errno"
	"errors"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var commentClient commentservice.Client

func initCommentRpc() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	c, err := commentservice.NewClient(
		"comment",
		client.WithMuxConnection(1),
		client.WithRPCTimeout(2*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	commentClient = c
}

func CreateComment(ctx context.Context, req *comment.CreateCommentRequest) (*common.Response, *common.Comment) {
	resp, err := commentClient.CreateComment(ctx, req)
	if err != nil {
		return &common.Response{
			StatusCode: errno.ServiceErrCode,
			StatusMsg:  err.Error(),
		}, nil
	}
	if resp.BaseResp.StatusCode != 0 {
		return &common.Response{
			StatusCode: resp.BaseResp.StatusCode,
			StatusMsg:  resp.BaseResp.StatusMessage,
		}, nil
	}
	return &common.Response{
			StatusCode: resp.BaseResp.StatusCode,
			StatusMsg:  resp.BaseResp.StatusMessage,
		},
		&common.Comment{
			Id: resp.Comment.CommentId,
			User: common.User{
				Id:            resp.Comment.UserId,
				Name:          "testName",
				FollowCount:   10,
				FollowerCount: 10,
				IsFollow:      false,
			},
			Content:    resp.Comment.Content,
			CreateDate: resp.Comment.CreateDate,
		}
}

func DeleteComment(ctx context.Context, req *comment.DeleteCommentRequest) *common.Response {
	resp, err := commentClient.DeleteComment(ctx, req)
	if err != nil {
		return &common.Response{
			StatusCode: errno.ServiceErrCode,
			StatusMsg:  err.Error(),
		}
	}
	if resp.BaseResp.StatusCode != 0 {
		return &common.Response{
			StatusCode: resp.BaseResp.StatusCode,
			StatusMsg:  resp.BaseResp.StatusMessage,
		}
	}
	return &common.Response{StatusCode: resp.BaseResp.StatusCode,
		StatusMsg: resp.BaseResp.StatusMessage,
	}
}

func QueryComments(ctx context.Context, req *comment.QueryCommentsRequest) (*common.Response, []*common.Comment) {
	resp, err := commentClient.QueryComments(ctx, req)
	if err != nil {
		return &common.Response{
			StatusCode: errno.ServiceErrCode,
			StatusMsg:  err.Error(),
		}, nil
	}
	if resp.BaseResp.StatusCode != 0 {
		return &common.Response{
			StatusCode: resp.BaseResp.StatusCode,
			StatusMsg:  resp.BaseResp.StatusMessage,
		}, nil
	}
	return &common.Response{
		StatusCode: resp.BaseResp.StatusCode,
		StatusMsg:  resp.BaseResp.StatusMessage,
	}, changeCommonComment(resp)
}

func changeCommonComment(resp *comment.QueryCommentsResponse) []*common.Comment {
	size := len(resp.Comments)
	res := make([]*common.Comment, size)

	for i := 0; i < size; i++ {
		res[i] = &common.Comment{
			Id: resp.Comments[i].CommentId,
			User: common.User{
				Id:            resp.Comments[i].UserId,
				Name:          "testName",
				FollowCount:   10,
				FollowerCount: 10,
				IsFollow:      false,
			},
			Content:    resp.Comments[i].Content,
			CreateDate: resp.Comments[i].CreateDate,
		}
	}

	return res
}

func QueryCommentNumber(ctx context.Context, req *comment.QueryCommentNumberRequest) (int64, error) {
	resp, err := commentClient.QueryCommentNumber(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errors.New(resp.BaseResp.StatusMessage)
	}
	return resp.CommentNumber, nil
}
