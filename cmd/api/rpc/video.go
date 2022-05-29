package rpc

import (
	"context"
	"time"

	"douyin/kitex_gen/video"
	"douyin/kitex_gen/video/videoservice"
	"douyin/pkg/constants"
	"douyin/pkg/errno"

	//"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/errno"
	//"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/middleware"

	"github.com/cloudwego/kitex/client"
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

func PublishVideo(ctx context.Context, req *video.PublishVideoRequest) error {
	resp, err := videoClient.PublishVideo(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		//TODO错误处理
		return errno.NewErrNo(resp.BaseResp.StatusCode, *(resp.BaseResp.StatusMsg))
	}
	return nil
}

// QueryNotes query list of note info
func GetPublishList(ctx context.Context, req *video.GetPublishListRequest) ([]*video.Video, error) {
	resp, err := videoClient.GetPublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		//TODO错误处理
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, *(resp.BaseResp.StatusMsg))
	}
	return resp.VideoList, nil
}

// UpdateNote update note info
func GetFeed(ctx context.Context, req *video.GetFeedRequest) ([]*video.Video, int64, error) {

	resp, err := videoClient.GetFeed(ctx, req)

	if err != nil {
		return nil, 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		//TODO错误处理
		return nil, 0, errno.NewErrNo(resp.BaseResp.StatusCode, *(resp.BaseResp.StatusMsg))
	}
	return resp.VideoList, *resp.NextTime, nil
}
