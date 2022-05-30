package handler

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/video"
	"douyin/pkg/errno"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PublishVideo(c *gin.Context) {

	var queryVar struct {
		Token string `json:"token" form:"token"`
		Data  []byte `json:"data" form:"data"`
		Title string `json:"title" form:"title"`
	}

	if err := c.BindQuery(&queryVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	//token处理
	userId, err := strconv.ParseInt(queryVar.Token, 10, 64)
	if err != nil {
		fmt.Println("handler.PublishVideo()->strconv.ParseInt error")
		return
	}
	//data文件处理
	//TODO错误处理
	//TODO错误处理
	if userId < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	req := &video.PublishVideoRequest{Token: queryVar.Token, Data: queryVar.Data, Title: queryVar.Title}

	err = rpc.PublishVideo(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

// func VerifyVideoId(c *gin.Context) {

// 	var queryVar struct {
// 		Token   string `json:"token" form:"token"`
// 		VideoId int64  `json:"videoId" form:"videoId"`
// 	}

// 	if err := c.BindQuery(&queryVar); err != nil {
// 		SendResponse(c, errno.ConvertErr(err), nil)
// 	}

// 	req := &video.VerifyVideoIdRequest{Token: queryVar.Token, VideoId: queryVar.VideoId}
// 	fmt.Println(req.VideoId)
// 	rOrf, err := rpc.VerifyVideoId(context.Background(), req)
// 	if err != nil {
// 		SendResponse(c, errno.ConvertErr(err), rOrf)
// 		return
// 	}
// 	SendResponse(c, errno.Success, rOrf)
// }
