package db

import (
	"context"
	"gorm.io/gorm"
)

// User user model
type User struct {
	gorm.Model
	Username    string `json:"u_name"`
	Password    string `json:"passwd"`
	FollowCount int64  `json:"follow_count"`
	FansCount   int64  `json:"fans_count"`
	Token       string `json:"token"`
}

// UserInfo user model
type UserInfo struct {
	ID          int64  `json:"u_id"`
	Username    string `json:"u_name"`
	Password    string `json:"passwd"`
	FollowCount int64  `json:"follow_count"`
	FansCount   int64  `json:"fans_count"`
}

type UserToken struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func (u *User) TableName() string {
	// TODO: add user table name here
	return "users"
}

// GetUserInfo  do db operation
func GetUserInfo(ctx context.Context, userID int64) (*UserInfo, error) {
	var user *UserInfo
	// TODO: db operation
	if err := DB.WithContext(ctx).Where("id == ?", userID).Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser do db operation
func CreateUser(ctx context.Context, username, password string) (*UserToken, error) {
	// TODO: db operation
	var res *UserToken
	err := DB.WithContext(ctx).Create(User{
		Username: username,
		Password: password,
	}).Error
	if err != nil {
		return nil, err
	}

	// TODO: get id and token
	return res, nil
}

// CheckUser do db operation
func CheckUser(ctx context.Context, username, password string) (*UserToken, error) {
	var res *UserToken
	if err := DB.WithContext(ctx).Where("user_name = ? and password = ?", username, password).Find(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
