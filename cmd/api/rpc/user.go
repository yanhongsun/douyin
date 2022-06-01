package rpc

import (
	"context"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/userservice"
	"douyin/middleware"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := userservice.NewClient(
		constants.UserServiceName,
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

// UserInfo user info format
type UserInfo struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func QueryUser(ctx context.Context, req *user.DouyinUserRequest) (*UserInfo, error) {
	resp, err := userClient.QueryCurUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}

	var userInfo UserInfo
	userInfo.ID = resp.User.Id
	userInfo.Name = resp.User.Name
	userInfo.FollowCount = resp.User.GetFollowCount()
	userInfo.FollowerCount = resp.User.GetFollowerCount()
	userInfo.IsFollow = resp.User.IsFollow
	return &userInfo, nil
}

// TODO: remove later
func IsUserExisted(ctx context.Context, req *user.DouyinUserExistRequest) (bool, error) {
	resp, err := userClient.IsUserExisted(ctx, req)
	if err != nil {
		return false, err
	}
	if resp.StatusCode != 0 {
		return false, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}

	return resp.IsExisted, nil
}

// TODO: remove later
func QueryOtherUser(ctx context.Context, req *user.DouyinQueryUserRequest) (*UserInfo, error) {
	resp, err := userClient.QueryOtherUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}

	var userInfo UserInfo
	userInfo.ID = resp.User.Id
	userInfo.Name = resp.User.Name
	userInfo.FollowCount = resp.User.GetFollowCount()
	userInfo.FollowerCount = resp.User.GetFollowerCount()
	userInfo.IsFollow = resp.User.IsFollow
	return &userInfo, nil
}

// TODO: remove later
func MultiQueryUser(ctx context.Context, req *user.DouyinMqueryUserRequest) ([]*user.User, error) {
	resp, err := userClient.MultiQueryUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}

	return resp.Users, nil
}
