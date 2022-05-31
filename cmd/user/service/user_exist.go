package service

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/user"
)

type UserExistService struct {
	ctx context.Context
}

func NewUserExistService(ctx context.Context) *UserExistService {
	return &UserExistService{
		ctx: ctx,
	}
}

// UserExist call db to check user is valid or not
func (s *UserExistService) UserExist(req *user.DouyinUserExistRequest) (bool, error) {
	targetID := req.TargetId
	res, err := db.QueryUserByID(s.ctx, targetID)
	if err != nil {
		return false, err
	}
	if len(res) == 0 {
		return false, nil
	}
	return true, nil
}
