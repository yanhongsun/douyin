package db

import (
	"douyin/pkg/constants"

	"context"
)

type User struct {
	Id            int64  `gorm:"column:id"`
	Name          string `gorm:"column:u_name"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:fans_count"`
}

func (v User) TableName() string {
	return constants.UserTableName
}

//查询用户信息 ok
func GetUserinfor(ctx context.Context, userId int64) (User, error) {
	var usr User
	if err := DB.WithContext(ctx).Where("id=?", userId).Take(&usr).Error; err != nil {
		return usr, err
	}
	return usr, nil
}

// func TestGetUserinfor(userId int64) (User, error) {
// 	var usr User
// 	if err := DB.Where("id=?", userId).Take(&usr).Error; err != nil {
// 		return usr, err
// 	}
// 	return usr, nil
// }
