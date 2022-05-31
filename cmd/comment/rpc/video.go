package rpc

import (
	"context"
	"time"

	"douyin/kitex_gen/video"
	"douyin/kitex_gen/video/videoservice"
	"douyin/pkg/constants"

	//"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/errno"
	//"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/middleware"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var videoClient videoservice.Client

func initVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		constants.VideoServiceName,
		//client.WithMiddleware(middleware.CommonMiddleware),
		//client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		//client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r), // resolver
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func VerifyVideoId(ctx context.Context, videoId int64, token string) (bool, error) {

	req := &video.VerifyVideoIdRequest{VideoId: videoId, Token: token}

	resp, err := videoClient.VerifyVideoId(ctx, req)

	klog.Fatal(resp)

	if err != nil {
		return false, err
	}

	return resp.TOrf, nil
}
