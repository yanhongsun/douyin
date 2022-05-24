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
	commentIndex := CommentIndex{VedioID: vedioId, CommentsNumber: 0}
	return DB.WithContext(ctx).Create(commentIndex).Error
}

func QueryCommentsNumber(ctx context.Context, vedioId int64) (int64, error) {
	var res CommentIndex
	if err := DB.WithContext(ctx).Model(&CommentIndex{}).Where("vedio_id = ?", vedioId).Find(&res).Error; err != nil {
		return res.CommentsNumber, err
	}
	return res.CommentsNumber, nil
}
