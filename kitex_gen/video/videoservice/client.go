// Code generated by Kitex v0.3.1. DO NOT EDIT.

package videoservice

import (
	"context"
	"douyin/kitex_gen/video"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	PublishVideo(ctx context.Context, req *video.PublishVideoRequest, callOptions ...callopt.Option) (r *video.PublishVideoResponse, err error)
	GetPublishList(ctx context.Context, req *video.GetPublishListRequest, callOptions ...callopt.Option) (r *video.GetPublishListResponse, err error)
	GetFeed(ctx context.Context, req *video.GetFeedRequest, callOptions ...callopt.Option) (r *video.GetFeedResponse, err error)
	VerifyVideoId(ctx context.Context, req *video.VerifyVideoIdResponse, callOptions ...callopt.Option) (r *video.VerifyVideoIdRequest, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kVideoServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kVideoServiceClient struct {
	*kClient
}

func (p *kVideoServiceClient) PublishVideo(ctx context.Context, req *video.PublishVideoRequest, callOptions ...callopt.Option) (r *video.PublishVideoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishVideo(ctx, req)
}

func (p *kVideoServiceClient) GetPublishList(ctx context.Context, req *video.GetPublishListRequest, callOptions ...callopt.Option) (r *video.GetPublishListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPublishList(ctx, req)
}

func (p *kVideoServiceClient) GetFeed(ctx context.Context, req *video.GetFeedRequest, callOptions ...callopt.Option) (r *video.GetFeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFeed(ctx, req)
}

func (p *kVideoServiceClient) VerifyVideoId(ctx context.Context, req *video.VerifyVideoIdResponse, callOptions ...callopt.Option) (r *video.VerifyVideoIdRequest, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VerifyVideoId(ctx, req)
}
