package assist

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/kitex_gen/video"
	"douyin/pkg/constants"
	"fmt"
	"time"
)

func GetBaseResp(stausCode int32, statusMsg string) video.BaseResp {
	return video.BaseResp{
		StatusCode: stausCode,
		StatusMsg:  &statusMsg,
	}
}

// 为db.Video类型变量赋值
func InitDBVideo(usrId int64, playUrl, coverUrl, title string) db.Video {
	return db.Video{
		UserId:        usrId,
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		Title:         title,
		FavoriteCount: 0,
		CommentCount:  0,
		CreateTime:    time.Now().Unix(),
	}

}

/*
	Id            int64  `gorm:"column:id"`
	UserId        int64  `gorm:"column:u_id"`
	PlayUrl       string `gorm:"column:play_url"`
	CoverUrl      string `gorm:"column:cover_url"`
	Title         string `gorm:"column:title"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	CreateTime    int64  `gorm:"column:create_time"`
*/
/*
type GetPublishListResponse struct {
	BaseResp  *BaseResp `thrift:"base_resp,1" json:"base_resp"`
	VideoList []*Video  `thrift:"video_list,2" json:"video_list"`
}
*/
func GetPublishListResp(ctx context.Context, videos []*db.Video, userId int64) video.GetPublishListResponse {
	var vid []*video.Video
	var bresp video.BaseResp
	if vid = GetVideoList(ctx, videos, userId); vid == nil {
		bresp = GetBaseResp(0, "发布列表获取成功！")
	} else {
		bresp = GetBaseResp(0, "无发布列表！")
	}

	return video.GetPublishListResponse{
		VideoList: vid,
		BaseResp:  &bresp,
	}

}

func GetFeedResp(ctx context.Context, videos []*db.Video, userId int64) video.GetFeedResponse {
	var vid []*video.Video
	var bresp video.BaseResp
	if vid = GetVideoList(ctx, videos, userId); vid == nil {
		bresp = GetBaseResp(0, "视频流获取成功！")
	} else {
		bresp = GetBaseResp(0, "无视频！")
	}
	return video.GetFeedResponse{
		VideoList: vid,
		BaseResp:  &bresp,
		NextTime:  &(*videos[len(videos)-1]).CreateTime,
	}

}

/*
type User struct {
	Id            int64  `thrift:"id,1,required" json:"id"`
	Name          string `thrift:"name,2,required" json:"name"`
	FollowCount   *int64 `thrift:"follow_count,3" json:"follow_count,omitempty"`
	FollowerCount *int64 `thrift:"follower_count,4" json:"follower_count,omitempty"`
	IsFollow      bool   `thrift:"is_follow,5,required" json:"is_follow"`
}

type User struct {
	gorm.Model
	Id            int64  `gorm:"column:id"`
	Name          string `gorm:"column:u_name"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:fans_count"`
}
*/
func foruser(ctx context.Context, u db.User, userId int64) video.User {

	var isf bool
	if userId == constants.EmptyUserId {
		isf = false
	} else {
		isf = db.IsFollow(ctx, userId, u.Id)
	}

	return video.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   &u.FollowCount,
		FollowerCount: &u.FollowerCount,
		IsFollow:      isf,
	}
}
func forvideo(ctx context.Context, v db.Video, userId int64) *video.Video {

	dbusr, err := db.GetUserinfor(ctx, v.UserId)
	var usr video.User

	if err == nil {
		usr = foruser(ctx, dbusr, userId)
	} else {
		//TODO错误处理，打印日志
		fmt.Println("assist.forvideo()-> db.GetUserinfor  error")
		return nil
	}
	isf := false
	if userId != constants.EmptyUserId {
		isf = db.IsFavorite(ctx, userId, v.Id)
	}
	return &video.Video{
		Id:            v.Id,
		Author:        &usr,
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		Title:         v.Title,
		IsFavorite:    isf,
	}
}
func GetVideoList(ctx context.Context, vid []*db.Video, userId int64) []*video.Video {

	vv := make([]*video.Video, 0)
	for _, v := range vid {
		if dv := forvideo(ctx, *v, userId); dv != nil {
			vv = append(vv, dv)
		}
	}
	return vv
}

/*

type Video struct {
	Id            int64  `thrift:"id,1,required" json:"id"`
	Author        *User  `thrift:"author,2,required" json:"author"`
	PlayUrl       string `thrift:"play_url,3,required" json:"play_url"`
	CoverUrl      string `thrift:"cover_url,4,required" json:"cover_url"`
	FavoriteCount int64  `thrift:"favorite_count,5,required" json:"favorite_count"`
	CommentCount  int64  `thrift:"comment_count,6,required" json:"comment_count"`
	IsFavorite    bool   `thrift:"is_favorite,7,required" json:"is_favorite"`
	Title         string `thrift:"title,8,required" json:"title"`
}
*/
