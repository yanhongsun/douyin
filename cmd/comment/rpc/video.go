package rpc

import (
	"context"
	"time"

	"douyin/cmd/comment/pack/configdata"
	"douyin/cmd/comment/pack/zapcomment"
	"douyin/kitex_gen/video"
	"douyin/kitex_gen/video/videoservice"

	//"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/errno"
	//"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/middleware"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var videoClient videoservice.Client

func initVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{configdata.CommentServerConfig.EtcdHost})
	if err != nil {
		zapcomment.Logger.Panic("etcd initialization err: " + err.Error())
	}

	c, err := videoservice.NewClient(
		configdata.CommentServerConfig.VideoServName,
		//client.WithMiddleware(middleware.CommonMiddleware),
		//client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		//client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithResolver(r), // resolver
	)
	if err != nil {
		zapcomment.Logger.Panic("videoService initialization err: " + err.Error())
	}
	videoClient = c
	zapcomment.Logger.Info("videoService initialization succeeded")
}

func VerifyVideoId(ctx context.Context, videoId int64) (bool, error) {

	req := &video.VerifyVideoIdRequest{VideoId: videoId, Token: ""}

	resp, err := videoClient.VerifyVideoId(ctx, req)

	if err != nil {
		return false, err
	}

	return resp.TOrf, nil
}
