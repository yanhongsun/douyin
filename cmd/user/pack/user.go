package pack

import (
	"github.com/douyin/cmd/user/dal/db"
	"github.com/douyin/kitex_gen/douyin_user"
)

func User(u *db.User) *douyin_user.User {
	if u == nil {
		return nil
	}

	return &douyin_user.User{
		Id:   int64(u.ID),
		Name: u.Username,
	}
}
