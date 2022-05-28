package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin/cmd/api/handlers"
	"github.com/douyin/cmd/api/rpc"
	"github.com/douyin/pkg/tracer"
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
	userGroup.POST("/", handlers.QueryUser)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
