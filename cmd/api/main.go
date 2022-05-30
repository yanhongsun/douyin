package main

import (
	"douyin/cmd/api/controller"
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/middleware"
	"douyin/cmd/api/rpc"
	"douyin/pkg/tracer"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	tracer.InitJaegers("douyin_api")
	rpc.InitRPC()

	r.Use(middleware.OpenTracing())

	// public directory is used to serve static resources
	r.Static("/resource", "./resource")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", handlers.GetFeed)
	apiRouter.GET("/user/", handlers.QueryUser)
	apiRouter.POST("/user/register/", handlers.Register)
	apiRouter.POST("/user/login/", handlers.Login)
	apiRouter.POST("/publish/action/", handlers.GetPublishList)
	apiRouter.GET("/publish/list/", handlers.GetPublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}

func main() {
	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
