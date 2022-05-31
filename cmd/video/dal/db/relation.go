package db

import (
	"douyin/pkg/constants"
	"errors"
	"fmt"

	"context"

	"gorm.io/gorm"
)

type Relation struct {
	Follower1 int64 `gorm:"column:Follower1"`
	Follower2 int64 `gorm:"column:Follower2"`
	Tag       int   `gorm:"column:tag"`
}

//为Video绑定表名
func (v Relation) TableName() string {
	return constants.RelationTableName
}

//ok
func IsFollow(ctx context.Context, meId int64, userId int64) bool {
	var rela Relation
	var big int64
	var small int64
	if meId > userId {
		big = meId
		small = userId
	} else {
		big = userId
		small = meId
	}
	err := DB.WithContext(ctx).Where("follower1=? and follower2=?", small, big).Take(&rela).Error
	if err == nil {
		if (meId == big && rela.Tag != 1) || (meId == small && rela.Tag != 2) {
			return true
		} else {
			return false
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	fmt.Println("db.IsFollow() heppen unknow error!", err)
	return false
}
