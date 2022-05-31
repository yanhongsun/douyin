package db

import (
	"douyin/pkg/constants"
	"errors"

	"context"

	"gorm.io/gorm"
)

type Video struct {
	Id            int64  `gorm:"column:id"`
	UserId        int64  `gorm:"column:u_id"`
	PlayUrl       string `gorm:"column:play_url"`
	CoverUrl      string `gorm:"column:cover_url"`
	Title         string `gorm:"column:title"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	CreateTime    int64  `gorm:"column:create_time"`
}

//为Video绑定表名
func (v Video) TableName() string {
	return constants.VideoTableName
}

//插入视频数据  ok
func PublishVideo(ctx context.Context, vid Video) error {
	//通过DB.WithContext(ctx) 对Context 支持，可为*gorm.DB 设置超时 Context等
	//select()
	if err := DB.WithContext(ctx).Create(&vid).Error; err != nil {
		return err
	}
	return nil
}

//查询用户发布列表
func GetPublishList(ctx context.Context, userId int64) ([]*Video, error) {
	//
	var vid []*Video
	err := DB.WithContext(ctx).Where("u_id=?", userId).Find(&vid).Error
	if err == nil {
		return vid, nil
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return nil, err
}

func GetFeed(ctx context.Context, lastTime int64, limit int) ([]*Video, error) {
	var vid []*Video
	if lastTime == 0 {
		lastTime = constants.MaxTime
	}
	err := DB.WithContext(ctx).Where("create_time <= ?", lastTime).Order("create_time desc").Limit(limit).Find(&vid).Error
	if err == nil {
		return vid, nil
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return nil, err
}

func VerifyVideoId(ctx context.Context, videoId int64) (bool, error) {
	var vid Video
	err := DB.WithContext(ctx).Where("id=?", videoId).Take(&vid).Error
	if err == nil {
		return true, nil
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return false, err
}
