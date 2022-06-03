package service

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/cmd/user/rpc"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
)

type QueryUserService struct {
	ctx context.Context
}

func NewQueryUserService(ctx context.Context) *QueryUserService {
	return &QueryUserService{
		ctx: ctx,
	}
}

// QueryCurUserByID Query user information by id
func (s *QueryUserService) QueryCurUserByID(req *user.DouyinUserRequest) (*user.User, error) {
	userInfo, err := db.QueryUserByID(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	if len(userInfo) == 0 {
		return nil, errno.UserNotExistErr
	}

	u := userInfo[0]
	return &user.User{
		Id:            u.ID,
		Name:          u.Username,
		FollowCount:   &u.FollowCount,
		FollowerCount: &u.FollowerCount,
		IsFollow:      false,
	}, nil
}

// QueryOtherUserByID Query user information by id
func (s *QueryUserService) QueryOtherUserByID(req *user.DouyinQueryUserRequest) (*user.User, error) {
	userInfo, err := db.QueryUserByID(s.ctx, req.TargetId)
	if err != nil {
		return nil, err
	}

	if len(userInfo) == 0 {
		return nil, errno.UserNotExistErr
	}

	u := userInfo[0]

	follow, err := rpc.IsFollow()

	return &user.User{
		Id:            u.ID,
		Name:          u.Username,
		FollowCount:   &u.FollowCount,
		FollowerCount: &u.FollowerCount,
		IsFollow:      false,
	}, nil
}
