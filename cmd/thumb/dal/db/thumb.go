package db

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type Video struct {
	Id            int64  `json:"id"`
	UId           int64  `gorm:"column:u_id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	Title         string `json:"title"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	CreateTime    string `json:"create_time"`
}
type Favorite struct {
	UId int64 `gorm:"column:u_id;"`
	VId int64 `gorm:"column:v_id;"`
}
type User struct {
	ID            int64  `json:"id"`
	Username      string `gorm:"column:u_name;"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `gorm:"column:fans_count"`
}

func UpdatdeVideo(ctx context.Context, uid, vid, actionType int64) error {
	fmt.Println("进入UpdatdeVideo,",uid,vid,actionType)
	//更新video表
	var expr string
	//点赞为1，取消点赞为-1
	if actionType == 1 {
		expr = "favorite_count+1"
	} else {
		expr = "favorite_count-1"
	}
	if err := DB.Model(&Video{}).Where("id = ? ", vid).Update("favorite_count", gorm.Expr(expr)).Error; err != nil {
		return fmt.Errorf("err in UpdatdeVideo when update_video_table:[%w]", err)
	}
	favorite := Favorite{UId: uid, VId: vid}
	fmt.Println("favorite:", favorite)
	//更新favorite表
	if actionType == 1 {
		if err := DB.WithContext(ctx).Create(&favorite).Error; err != nil {
			return fmt.Errorf("err in UpdatdeVideo when update_favorite_table:[%w]", err)
		}
	} else {
		if err := DB.WithContext(ctx).Where("u_id = ? and v_id = ?", uid, vid).Delete(&favorite).Error; err != nil {
			return fmt.Errorf("err in UpdatdeVideo when update_favorite_table:[%w]", err)
		}
	}

	return nil
}

func ListVideo(ctx context.Context, uid int64) ([]*Video, error) {
	var total int64
	var res []*Favorite
	var videos []*Video
	conn := DB.WithContext(ctx).Model(&Favorite{}).Where("u_id = ?", uid).Find(&res)
	fmt.Println("ListVideo中 res:", res)

	for i, x := range res {
		fmt.Println(i,x )
	}
	//没用？？
	if err := conn.Count(&total).Error; err != nil {
		return videos, err
	}
	//获得所有vid
	if err := conn.Find(&res).Error; err != nil {
		return videos, err
	}	
	vIDs := []int64{}
	for _, fav := range res {
		vIDs = append(vIDs, fav.VId)
	}
	fmt.Println("vIDs:", vIDs)
	//得到所有video信息
	if err := DB.WithContext(ctx).Where("id in ?", vIDs).Find(&videos).Error; err != nil {
		return videos, err
	}

	fmt.Println("ListVideo中 videos:", videos)
	for i, x := range videos {
		fmt.Println(i,x )
	}
	return videos, nil
}

func GetUserInfo(ctx context.Context, userID int64) (*User, error) {
	fmt.Println("GetUserInfo userID:", userID)
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("id = ?", userID).Find(&res).Error; err != nil {
		return nil, err
	}
	fmt.Println("GetUserInfo:")
	for i, x := range res {
		fmt.Println(i,x )
	}
	u := res[0]
	return &User{
		ID:            u.ID,
		Username:      u.Username,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
	}, nil
}
