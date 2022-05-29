// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"github.com/yanhongsun/douyin/cmd/api/handlers"
	"github.com/yanhongsun/douyin/cmd/api/rpc"
	"github.com/yanhongsun/douyin/pkg/constants"
	"github.com/yanhongsun/douyin/pkg/tracer"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.New()
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		//Authenticator: func(c *gin.Context) (interface{}, error) {
		//	var loginVar handlers.UserParam
		//	if err := c.ShouldBind(&loginVar); err != nil {
		//		return "", jwt.ErrMissingLoginValues
		//	}
		//
		//	if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
		//		return "", jwt.ErrMissingLoginValues
		//	}
		//
		//	return rpc.CheckUser(context.Background(), &like.CheckUserRequest{UserName: loginVar.UserName, Password: loginVar.PassWord})
		//},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	//v1 := r.Group("/v1")
	//user1 := v1.Group("/user")
	//user1.POST("/login", authMiddleware.LoginHandler)
	//user1.POST("/register", handlers.Register)

	note1 := r.Group("/note")
	note1.Use(authMiddleware.MiddlewareFunc())
	note1.GET("/thumblist", handlers.ThumbList)
	note1.POST("/action", handlers.Likeyou)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
