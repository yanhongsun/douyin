package db

import (
	"context"
	"douyin/pkg/constants"
)

// User user model
type User struct {
	ID       int64  `gorm:"column:id;primaryKey;not null"`
	Username string `gorm:"column:u_name;unique;type:varchar(30);not null"`
	Password string `gorm:"column:passwd;type:varchar(32);not null"`
	// Nickname    string `json:"nickname"`
	FollowCount   int64 `gorm:"column:follow_count;default:0"`
	FollowerCount int64 `gorm:"column:fans_count;default:0"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// QueryUserByName get user information by username
func QueryUserByName(ctx context.Context, username string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("u_name = ?", username).Limit(1).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// QueryUserByID get user information by id
func QueryUserByID(ctx context.Context, targetID int64) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("id = ?", targetID).Limit(1).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// MultiQueryUserByID get multiple users by user id
func MultiQueryUserByID(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser do db operation
func CreateUser(ctx context.Context, username, password string) ([]*User, error) {
	users := []*User{{
		Username: username,
		Password: password,
	}}
	res := make([]*User, 0)

	err := DB.WithContext(ctx).Create(users).Error
	if err != nil {
		return nil, err
	}
	// get token and id
	err = DB.WithContext(ctx).Where("u_name = ?", username).Limit(1).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateSalt(ctx context.Context, username string, salt []byte) error {
	return nil
}

func QuerySalt(ctx context.Context, username string) ([]byte, error) {
	return nil, nil
}
