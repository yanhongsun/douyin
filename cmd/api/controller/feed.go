package controller

import (
	"douyin/cmd/api/common"
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
	c.JSON(http.StatusOK, FeedResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
