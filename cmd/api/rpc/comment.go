package rpc

import (
	"context"
	"douyin/cmd/api/common"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/comment/commentservice"
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

func CreateComment(ctx context.Context, req *comment.CreateCommentRequest) (*common.Comment, error) {
	resp, err := commentClient.CreateComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errors.New("something wrong")
	}
	return &common.Comment{
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
	}, nil
}
