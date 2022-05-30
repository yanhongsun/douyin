package rpc

import (
	"context"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/comment/commentservice"
	"douyin/pkg/errno"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
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
		client.WithSuite(trace.NewDefaultClientSuite()),
	)
	if err != nil {
		panic(err)
	}
	commentClient = c
}

type Response struct {
	StatusCode int32
	StatusMsg  string
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

func CreateComment(ctx context.Context, req *comment.CreateCommentRequest) (*Response, *Comment) {
	resp, err := commentClient.CreateComment(ctx, req)
	if err != nil {
		return &Response{
			StatusCode: errno.ServiceErrCode,
			StatusMsg:  err.Error(),
		}, nil
	}
	if resp.BaseResp.StatusCode != 0 {
		return &Response{
			StatusCode: resp.BaseResp.StatusCode,
			StatusMsg:  resp.BaseResp.StatusMessage,
		}, nil
	}
	return &Response{
			StatusCode: resp.BaseResp.StatusCode,
			StatusMsg:  resp.BaseResp.StatusMessage,
		},
		&Comment{
			Id: resp.Comment.CommentId,
			User: User{
				Id:            resp.Comment.User.Id,
				Name:          resp.Comment.User.Name,
				FollowCount:   *resp.Comment.User.FollowCount,
				FollowerCount: *resp.Comment.User.FollowerCount,
				IsFollow:      resp.Comment.User.IsFollow,
			},
			Content:    resp.Comment.Content,
			CreateDate: resp.Comment.CreateDate,
		}
}

func DeleteComment(ctx context.Context, req *comment.DeleteCommentRequest) *Response {
	resp, err := commentClient.DeleteComment(ctx, req)
	if err != nil {
		return &Response{
			StatusCode: errno.ServiceErrCode,
			StatusMsg:  err.Error(),
		}
	}
	if resp.BaseResp.StatusCode != 0 {
		return &Response{
			StatusCode: resp.BaseResp.StatusCode,
			StatusMsg:  resp.BaseResp.StatusMessage,
		}
	}
	return &Response{StatusCode: resp.BaseResp.StatusCode,
		StatusMsg: resp.BaseResp.StatusMessage,
	}
}

func QueryComments(ctx context.Context, req *comment.QueryCommentsRequest) (*Response, []*Comment) {
	resp, err := commentClient.QueryComments(ctx, req)
	if err != nil {
		return &Response{
			StatusCode: errno.ServiceErrCode,
			StatusMsg:  err.Error(),
		}, nil
	}
	if resp.BaseResp.StatusCode != 0 {
		return &Response{
			StatusCode: resp.BaseResp.StatusCode,
			StatusMsg:  resp.BaseResp.StatusMessage,
		}, nil
	}
	return &Response{
		StatusCode: resp.BaseResp.StatusCode,
		StatusMsg:  resp.BaseResp.StatusMessage,
	}, changeCommonComment(resp)
}

func changeCommonComment(resp *comment.QueryCommentsResponse) []*Comment {
	size := len(resp.Comments)
	res := make([]*Comment, size)

	for i := 0; i < size; i++ {
		res[i] = &Comment{
			Id: resp.Comments[i].CommentId,
			User: User{
				Id:            resp.Comments[i].User.Id,
				Name:          resp.Comments[i].User.Name,
				FollowCount:   *resp.Comments[i].User.FollowCount,
				FollowerCount: *resp.Comments[i].User.FollowerCount,
				IsFollow:      resp.Comments[i].User.IsFollow,
			},
			Content:    resp.Comments[i].Content,
			CreateDate: resp.Comments[i].CreateDate,
		}
	}

	return res
}

// func QueryCommentNumber(ctx context.Context, req *comment.QueryCommentNumberRequest) (int64, error) {
// 	resp, err := commentClient.QueryCommentNumber(ctx, req)
// 	if err != nil {
// 		return 0, err
// 	}
// 	if resp.BaseResp.StatusCode != 0 {
// 		return 0, errors.New(resp.BaseResp.StatusMessage)
// 	}
// 	return resp.CommentNumber, nil
// }
