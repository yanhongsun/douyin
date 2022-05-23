package db

import (
	"context"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	CommentID int64  `json:"comment_id" gorm:"index:,sort:desc"`
	VedioID   int64  `json:"vedio_id" gorm:"index"`
	UserID    int64  `json:"user_id"`
	State     bool   `json:"state"`
	Content   string `json:"content"`
}

func CreateComment(ctx context.Context, comment Comment) error {
	return DB.WithContext(ctx).Create(comment).Error
}

func DeleteComment(ctx context.Context, commentId int64) error {
	return DB.WithContext(ctx).Where("comment_id = ?", commentId).Delete(&Comment{}).Error
}

func QueryComment(ctx context.Context, vedioId int64, limit, offset int) ([]*Comment, error) {
	var res []*Comment
	if err := DB.WithContext(ctx).Model(&Comment{}).Where("vedio_id = ?", vedioId).Limit(limit).Offset(offset).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
