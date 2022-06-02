package service

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/user"
)

type MultiQueryUserService struct {
	ctx context.Context
}

func NewMultiQueryUserService(ctx context.Context) *MultiQueryUserService {
	return &MultiQueryUserService{ctx: ctx}
}

func (s *MultiQueryUserService) MultiQueryUser(req *user.DouyinMqueryUserRequest) ([]*user.User, error) {
	userInfos := make([]*user.User, 0)
	users, err := db.MultiQueryUserByID(s.ctx, req.TargetIds)
	if err != nil {
		return nil, err
	}

	for _, u := range users {
		userInfos = append(userInfos, &user.User{
			Id:            u.ID,
			Name:          u.Username,
			FollowCount:   &u.FollowCount,
			FollowerCount: &u.FollowerCount,
			IsFollow:      false,
		})
	}
	/*for _, u := range users {
		isFollowed, err := rpc.IsFollow(s.ctx, &relation.IsFollowRequest{
			UserId:   req.UserId,
			ToUserId: u.ID,
			Token:    req.Token,
		})
		if err != nil {
			return nil, err
		}
		userInfos = append(userInfos, &user.User{
			Id:            u.ID,
			Name:          u.Username,
			FollowCount:   &u.FollowCount,
			FollowerCount: &u.FollowerCount,
			IsFollow:      isFollowed,
		})
	}*/
	return userInfos, nil
}
