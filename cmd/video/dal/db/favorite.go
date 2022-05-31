package db

import (
	"douyin/pkg/constants"
	"errors"
	"fmt"

	"context"

	"gorm.io/gorm"
)

type Favorite struct {
	Uid int64 `gorm:"column:u_id"`
	Vid int64 `gorm:"column:v_id"`
}

func (f Favorite) TableName() string {
	return constants.FavoriteTableName
}

//ok
func IsFavorite(ctx context.Context, userId int64, videoId int64) bool {
	var fav Favorite
	err := DB.WithContext(ctx).Where("u_id=? and v_id=?", userId, videoId).Take(&fav).Error
	if err == nil {
		return true
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	//TODO错误处理
	fmt.Println("db.IsFavorite() unknow error! ", err)
	return false
}
