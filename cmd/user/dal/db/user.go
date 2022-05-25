package db

import (
	"context"
	"gorm.io/gorm"
)

// TODO: modify
// User user model
type User struct {
	gorm.Model
	Username    string `json:"u_name"`
	Password    string `json:"passwd"`
	FollowCount int64  `json:"follow_count"`
	FansCount   int64  `json:"fans_count"`
	Token       string `json:"token"`
}

// TODO: modify
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

// QueryUser OK
func QueryUser(ctx context.Context, username string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("u_name = ?", username).Find(&res).Error; err != nil {
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
	err = DB.WithContext(ctx).Where("user_name = ?", username).Find(&res).Error
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
