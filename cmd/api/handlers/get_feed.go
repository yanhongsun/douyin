package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/video"
	"douyin/pkg/errno"

	"github.com/gin-gonic/gin"
)

func GetFeed(c *gin.Context) {

	var queryVar struct {
		Token      string `json:"token" form:"token"`
		LatestTime int64  `json:"latest_time,omitempty" form:"latest_time"`
	}

	//TODO错误处理
	if err := c.BindQuery(&queryVar); err != nil {

		SendResponseFeed(c, errno.ConvertErr(err), nil, 0)
	}

	if queryVar.LatestTime <= 0 {
		queryVar.LatestTime = 0
	}

	req := &video.GetFeedRequest{Token: &queryVar.Token, LatestTime: &queryVar.LatestTime}

	videoList, nextTime, err := rpc.GetFeed(context.Background(), req)

	if err != nil {
		SendResponseFeed(c, errno.ConvertErr(err), nil, 0)
		return
	}
	SendResponseFeed(c, errno.Success, videoList, nextTime)

}
