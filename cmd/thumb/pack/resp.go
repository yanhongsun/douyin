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

package pack

import (
	"douyin/kitex_gen/like"
	"douyin/pkg/errno"
	"errors"
)

// BuildLikeyouResp build LikeyouResponse from error and status
func BuildLikeyouResp(err error) *like.LikeyouResponse {
	if err == nil {
		return &like.LikeyouResponse{
			StatusCode: errno.Success.ErrCode,
			StatusMsg:  &errno.Success.ErrMsg,
		}
	}

	e := errno.ErrNo{}
	if errors.As(err, e) {
		// 不能：msg := err.(e).ErrMsg，因为e这里必须是个类型
		//todo:为啥不能直接取地址？ 应该是因为断言返回的是两个值，没有捕获这个err所以没法取地址
		msg := err.(errno.ErrNo).ErrMsg
		return &like.LikeyouResponse{
			StatusCode: err.(errno.ErrNo).ErrCode,
			StatusMsg:  &msg,
		}
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return &like.LikeyouResponse{
		StatusCode: errno.ServiceErrCode,
		StatusMsg:  &s.ErrMsg,
	}
}

func BuildThumblistResp(videoList []*like.Video, err error) *like.ThumbListResponse {
	if err == nil {
		return &like.ThumbListResponse{
			StatusCode: errno.SuccessCode,
			StatusMsg:  &errno.Success.ErrMsg,
			VideoList:  videoList,
		}
	}
	//不能带&！！
	if errors.As(err, errno.ErrNo{}) {
		msg := err.(errno.ErrNo).ErrMsg
		return &like.ThumbListResponse{
			StatusCode: errno.ServiceErrCode,
			StatusMsg:  &msg,
			VideoList:  nil,
		}
	}
	//为啥不能直接取地址？
	msg := err.Error()
	return &like.ThumbListResponse{
		StatusCode: errno.ServiceErrCode,
		StatusMsg:  &msg,
		VideoList:  videoList,
	}
}
