package rpc

import (
	"context"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/relation/relationservice"
	"douyin/pkg/errno"
	"fmt"
)

var relationClient relationservice.Client

func IsFollow(ctx context.Context, req *relation.IsFollowRequest) (bool, error) {
	// var resp relation.IsFollowResponse
	fmt.Println("=========== getIsFollow")
	resp, err := relationClient.IsFollow(ctx, req)
	if err != nil {
		return false, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return false, errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatusMsg)
	}
	return resp.IsFollow, nil
}
