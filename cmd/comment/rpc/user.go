package rpc

import (
	"context"
	"douyin/cmd/comment/pack/configdata"
	"douyin/cmd/comment/pack/zapcomment"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/userservice"
	"douyin/middleware"
	"douyin/pkg/errno"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{configdata.CommentServerConfig.EtcdHost})
	if err != nil {
		zapcomment.Logger.Panic("etcd initialization err: " + err.Error())
	}
	c, err := userservice.NewClient(
		configdata.CommentServerConfig.UserServName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		//client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithResolver(r), // resolver
	)
	if err != nil {
		zapcomment.Logger.Panic("userService initialization err: " + err.Error())
	}
	userClient = c
	zapcomment.Logger.Info("userService initialization succeeded")
}

type UserInfo struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func GetUserInfo(ctx context.Context, userId int64) (*UserInfo, error) {

	req := &user.DouyinUserRequest{UserId: userId, Token: ""}

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
