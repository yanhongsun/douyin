package pack

import (
	"github.com/douyin/cmd/user/dal/db"
	"github.com/douyin/kitex_gen/user"
)

func UserInfo(u *db.UserInfo) *user.User {
	if u == nil {
		return nil
	}

	var userInfo *user.User
	userInfo.SetId(int64(u.ID))
	userInfo.SetName(u.Username)
	userInfo.SetFollowCount(&u.FollowCount)
	userInfo.SetFollowerCount(&u.FansCount)

	return userInfo
}

func UserToken(u *db.UserToken) (int64, string) {
	return u.UserID, u.Token
}
