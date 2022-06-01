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
	"douyin/cmd/thumb/dal/db"
	"douyin/kitex_gen/like"
)

func User(u *db.User) *like.User {
	if u != nil {
		return nil
	}
	return &like.User{
		Id:            u.ID,
		Name:          u.Username,
		FollowCount:   &u.FollowCount,
		FollowerCount: &u.FollowerCount,
		IsFollow:      true, //没用的字段？
	}
}

//把db层的video转换成service层的
//为什么都是指针？
func Video(v *db.Video, u *db.User) *like.Video {
	if v != nil {
		return nil
	}
	return &like.Video{
		UserId:        u.ID,
		Author:        User(u),
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    true,
		Title:         v.Title,
	}
}

func Videos(v []*db.Video, u *db.User) []*like.Video {
	res := []*like.Video{}
	for _, video := range v {
		res = append(res, Video(video, u))
	}
	return res
}
