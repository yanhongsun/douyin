package service

import (
	"context"
	"douyin/cmd/user/dal/db"
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

	// TODO: get follow info from service relation
	var follow = false
	/*follow, err = rpc.IsFollowed(s.ctx, &relation.IsFollowRequest{
		UserId:   req.UserId,
		Token:    req.Token,
		ToUserId: req.TargetId,
	})*/

	u := userInfo[0]

	return &user.User{
		Id:            u.ID,
		Name:          u.Username,
		FollowCount:   &u.FollowCount,
		FollowerCount: &u.FollowerCount,
		IsFollow:      follow,
	}, nil
}
