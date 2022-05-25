package db

import (
	"context"

	"gorm.io/gorm"
)

type CommentIndex struct {
	gorm.Model
	VedioID        int64 `json:"vedio_id" gorm:"index"`
	CommentsNumber int64 `json:"comments_number"`
}

func CreateCommentIndex(ctx context.Context, vedioId int64) error {
	var count int64
	if err := DB.WithContext(ctx).Model(&CommentIndex{}).Where("vedio_id = ?", vedioId).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	return DB.WithContext(ctx).Create(&CommentIndex{VedioID: vedioId, CommentsNumber: 0}).Error
}

func QueryCommentsNumber(ctx context.Context, vedioId int64) (int64, error) {
	var res CommentIndex
	if err := DB.WithContext(ctx).Model(&CommentIndex{}).Where("vedio_id = ?", vedioId).Find(&res).Error; err != nil {
		return 0, err
	}
	return res.CommentsNumber, nil
}
