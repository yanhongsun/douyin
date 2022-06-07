package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/video"
	"douyin/pkg/errno"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func PublishVideo(c *gin.Context) {

	fmt.Println("api_PublishVideo")
	fmt.Println("Publish Video")
	var queryVar struct {
		Token string `json:"token" form:"token"`
		Data  []byte `json:"data" form:"data"`
		Title string `json:"title" form:"title"`
	}
	queryVar.Token = c.PostForm("token")
	queryVar.Title = c.PostForm("title")
	datafile, _ := c.FormFile("data")
	// if err := c.ShouldBind(&queryVar); err != nil {
	// 	SendResponseV(c, errno.ConvertErr(err), nil)
	// }

	data, err := datafile.Open()
	if err != nil {
		c.String(400, "文件格式错误")
		return
	}
	defer data.Close()
	queryVar.Data, err = ioutil.ReadAll(data)
	if err != nil {
		c.String(400, "文件错误")
		return
	}

	req := &video.PublishVideoRequest{Token: queryVar.Token, Data: queryVar.Data, Title: queryVar.Title}

	err = rpc.PublishVideo(context.Background(), req)
	if err != nil {
		SendResponseV(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponseV(c, errno.Success, nil)
}
