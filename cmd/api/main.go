package main

import (
	"github.com/douyin/cmd/api/handlers"
	"github.com/gin-gonic/gin"
)

func Init() {

}

func main() {
	Init()
	r := gin.New()

	douyin := r.Group("/douyin")

	userGroup := douyin.Group("/user")
	userGroup.POST("/login", handlers.Login) // TODO: login handler
	userGroup.POST("/register", handlers.Register)
	userGroup.GET("/", handlers.QueryUser) // TODO: get user information handler
}
