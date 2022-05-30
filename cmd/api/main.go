package main

import (
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/pkg/tracer"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	tracer.InitJaeger("api")
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.New()

	douyin := r.Group("/douyin")

	userGroup := douyin.Group("/user")
	userGroup.POST("/login", handlers.Login)
	userGroup.POST("/register", handlers.Register)
	userGroup.GET("/", handlers.QueryUser)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
