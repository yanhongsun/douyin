package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/douyin/cmd/api/handlers"
	"github.com/douyin/kitex_gen/user"
	"github.com/douyin/kitex_gen/user/userservice"
	"github.com/douyin/middleware"
	"github.com/douyin/pkg/errno"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var userClient userservice.Client

func initUserRpc() {
	// TODO: modify configs
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}
	c, err := userservice.NewClient(
		"user", // TODO: modify
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func CreateUser(ctx context.Context, req *user.DouyinUserRegisterRequest) (int64, string, error) {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return -1, "", err
	}

	if resp.StatusCode != 0 {
		return -1, "", errno.NewErrNo(resp.StatusCode, resp.GetStatusMsg())
	}

	return resp.UserId, resp.Token, nil
}

func CheckUser(ctx context.Context, req *user.DouyinUserLoginRequest) (int64, string, error) {
	// return id token error
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return -1, "", err
	}

	if resp.StatusCode != 0 {
		return -1, "", errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}

	return resp.UserId, resp.Token, nil
}

func GetUserInfo(ctx context.Context, req *user.DouyinUserRequest) (*handlers.UserInfo, error) {
	resp, err := userClient.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}

	var userInfo handlers.UserInfo
	userInfo.ID = resp.User.Id
	userInfo.Name = resp.User.Name
	userInfo.FollowCount = resp.User.GetFollowCount()
	userInfo.FollowerCount = resp.User.GetFollowerCount()
	userInfo.IsFollow = resp.User.IsFollow
	return &userInfo, nil
}
