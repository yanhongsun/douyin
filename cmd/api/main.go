package main

import (
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/pkg/tracer"
	"net/http"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func Init() {
	tracer.InitJaeger("api")
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.New()
	r.Static("/resource", "./resource")
	douyin := r.Group("/douyin")

	douyin.GET("/feed/", handlers.GetFeed)
	douyin.GET("/publish/list/", handlers.GetPublishList)
	douyin.POST("/publish/action/", handlers.PublishVideo)
	//vid.GET("/verifyVideoId/", handlers.VerifyVideoId)
	userGroup := douyin.Group("/user")
	userGroup.POST("/login/", handlers.Login)
	userGroup.POST("/register/", handlers.Register)
	userGroup.GET("/", handlers.QueryUser)

	commentGroup := douyin.Group("/comment")
	commentGroup.POST("/action/", handlers.CommentAction)
	commentGroup.GET("/list/", handlers.CommentList)

	if err := http.ListenAndServe(":8086", r); err != nil {
		klog.Fatal(err)
	}
}
