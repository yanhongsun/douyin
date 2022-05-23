package db

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	// TODO: use gorm or not?
	gorm.Model
	Username    string `json:"u_name"`
	Password    string `json:"passwd"`
	FollowCount int64  `json:"follow_count"`
	FansCount   int64  `json:"fans_count"`
	// TODO: add token or not?
	Token string `json:"token"`
}

func (u *User) TableName() string {
	// TODO: add user table name here
	return "users"
}

// GetUserInfo get information of specific user
func GetUserInfo(ctx context.Context, userID int64) (*User, error) {
	var user *User

	if err := DB.WithContext(ctx).Where("id == ?", userID).Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// UserRegister register an user
func UserRegister(ctx context.Context, username, password string) error {
	return DB.WithContext(ctx).Create(User{
		Username: username,
		Password: password,
	}).Error
}

// UserLogin user login app
func UserLogin(ctx context.Context, username, password string) (*User, error) {
	var res *User
	if err := DB.WithContext(ctx).Where("user_name = ? and password = ?", username, password).Find(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
