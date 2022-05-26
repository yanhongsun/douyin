package controller

import (
	"context"
	"douyin/cmd/api/common"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/comment"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	res, err := rpc.QueryCommentNumber(context.Background(), &comment.QueryCommentNumberRequest{VideoId: DemoVideos[0].Id})

	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: common.Response{StatusCode: 40007},
		})
		return
	}

	DemoVideos[0].CommentCount = res

	c.JSON(http.StatusOK, FeedResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
