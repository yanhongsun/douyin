package main

import (
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/middleware"
	"douyin/cmd/api/rpc"
	"net/http"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func Init() {
	// TODO
	// tracer.InitJaeger("api")
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
	userGroup.GET("/", middleware.AuthMiddleware(), handlers.QueryCurUser)
	// TODO: remove the handler for checking user existence
	userGroup.GET("/exist/", handlers.IsUserExisted)
	userGroup.GET("/otheruser/", handlers.QueryOtherUser)

	if err := http.ListenAndServe(":8090", r); err != nil {
		klog.Fatal(err)
	}
}
