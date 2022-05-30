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
	"net/http"

	"douyin/cmd/api/handler"
	"douyin/cmd/api/rpc"

	"github.com/cloudwego/kitex/pkg/klog"

	//"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/tracer"
	"github.com/gin-gonic/gin"
)

func Init() {
	// tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.New()
	r.Static("/resource", "./resource")
	vid := r.Group("/douyin")

	vid.GET("/feed/", handler.GetFeed)
	vid.GET("/publish/list/", handler.GetPublishList)
	vid.POST("/publish/action/", handler.PublishVideo)
	//vid.GET("/verifyVideoId/", handler.VerifyVideoId)
	if err := http.ListenAndServe(":8086", r); err != nil {
		klog.Fatal(err)
	}
}
