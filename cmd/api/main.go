package main

import (
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/middleware"
	"douyin/cmd/api/rpc"
	"douyin/pkg/tracer"
	"net/http"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func Init() {
	// TODO
	tracer.InitJaeger("api")
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.New()
	r.Use(middleware.OpenTracing())

	douyin := r.Group("/douyin")
	relationGroup := douyin.Group("/relation")
	relationGroup.GET("/follower/list/", handlers.GetFollowerList)
	relationGroup.GET("/follow/list/", handlers.GetFollowList)
	relationGroup.GET("/isfollow/", handlers.IsFollow)
	relationGroup.POST("/action/", handlers.RelationAction)

	r.Static("/resource", "./resource")

	douyin.GET("/feed/", handlers.GetFeed)
	douyin.GET("/publish/list/", handlers.GetPublishList)
	douyin.POST("/publish/action/", handlers.PublishVideo)
	//vid.GET("/verifyVideoId/", handlers.VerifyVideoId)
	userGroup := douyin.Group("/user")
	userGroup.POST("/login/", handlers.Login)
	userGroup.POST("/register/", handlers.Register)
	userGroup.GET("/", middleware.AuthMiddleware(), handlers.QueryCurUser)

	commentGroup := douyin.Group("/comment")
	commentGroup.POST("/action/", handlers.CommentAction)
	commentGroup.GET("/list/", handlers.CommentList)
	// test
	// userGroup.GET("/other/", handlers.QueryOthUser)
	// userGroup.GET("/mother/", handlers.MQueryUser)

	if err := http.ListenAndServe(":8086", r); err != nil {
		klog.Fatal(err)
	}
}
