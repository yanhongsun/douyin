package db

import (
	"context"
	"errors"

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
	return DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&CommentIndex{}).Where("vedio_id = ?", comment.VedioID).Update("comments_number", gorm.Expr("comments_number + ?", 1)).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Create(comment).Error; err != nil {
			return err
		}

		return nil
	})
}

func DeleteComment(ctx context.Context, commentId, vedioId int64) error {
	return DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var tmp CommentIndex
		if err := tx.Model(&CommentIndex{}).Where("vedio_id = ?", vedioId).Find(&tmp).Error; err != nil {
			return err
		}

		if tmp.CommentsNumber <= 0 {
			return errors.New("CommentsNumber is already 0")
		}

		if err := tx.Model(&CommentIndex{}).Where("vedio_id = ?", vedioId).Update("comments_number", gorm.Expr("comments_number + ?", -1)).Error; err != nil {
			return err
		}

		if err := tx.Where("comment_id = ?", commentId).Delete(&Comment{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func QueryComment(ctx context.Context, vedioId int64, limit, offset int) ([]*Comment, error) {
	var res []*Comment
	if err := DB.WithContext(ctx).Model(&Comment{}).Where("vedio_id = ?", vedioId).Limit(limit).Offset(offset).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
