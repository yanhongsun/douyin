package main

import (
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"net/http"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func Init() {
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.New()
	douyin := r.Group("/douyin")
	douyin.GET("/follower/list/", handlers.GetFollowerList)
	douyin.GET("/follow/list/", handlers.GetFollowList)
	douyin.POST("/action/", handlers.RelationAction)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
