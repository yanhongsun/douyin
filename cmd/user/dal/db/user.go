package db

import (
	"context"
	"github.com/douyin/kitex_gen/user"
)

// User user model
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"u_name"`
	Password string `json:"passwd"`
	// Nickname    string `json:"nickname"`
	FollowCount   int64 `json:"follow_count"`
	FollowerCount int64 `json:"fans_count"`
}

func (u *User) TableName() string {
	// TODO: add user table name here
	return "users"
}

// GetUserInfo  do db operation
func GetUserInfo(ctx context.Context, userID int64) (*user.User, error) {
	// TODO: db operation
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("id == ?", userID).Find(res).Error; err != nil {
		return nil, err
	}
	u := res[0]
	return &user.User{
		Id:            u.ID,
		Name:          u.Username,
		FollowCount:   &u.FollowCount,
		FollowerCount: &u.FollowerCount,
		IsFollow:      false,
	}, nil
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
