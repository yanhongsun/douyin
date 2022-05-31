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
	douyin.GET("/relation/follower/list/", handlers.GetFollowerList)
	douyin.GET("/relation/follow/list/", handlers.GetFollowList)
	douyin.GET("/relation/isfollow/", handlers.IsFollow)
	douyin.POST("/relation/action/", handlers.RelationAction)

	if err := http.ListenAndServe(":8081", r); err != nil {
		klog.Fatal(err)
	}
}
