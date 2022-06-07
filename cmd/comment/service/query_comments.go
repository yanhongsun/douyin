package service

import (
	"context"
	"douyin/cmd/comment/dal/mysqldb"
	"douyin/cmd/comment/dal/redisdb"
	"douyin/cmd/comment/pack"
	"douyin/cmd/comment/pack/zapcomment"
	"douyin/cmd/comment/repository"
	"douyin/kitex_gen/comment"
	"strconv"

	"golang.org/x/sync/singleflight"
)

var gQueryComments singleflight.Group

type QueryCommentsService struct {
	ctx context.Context
}

func NewQueryCommentsService(ctx context.Context) *QueryCommentsService {
	return &QueryCommentsService{ctx: ctx}
}

func (s *QueryCommentsService) QueryComments(req *comment.QueryCommentsRequest) ([]*comment.Comment, error) {
	status, res, err := redisdb.CheckCommentsCache(s.ctx, req.VideoId)

	if err != nil {
		zapcomment.Logger.Error("redisdb err: " + err.Error())
	}

	if status {
		resp, err := pack.ChangeComments(s.ctx, res.Comments)
		if err == nil {
			return resp, nil
		}
		zapcomment.Logger.Error("redisdb err: " + err.Error())
	}

	key := strconv.FormatInt(req.VideoId, 10)

	v, err, _ := gQueryComments.Do(key, func() (interface{}, error) {
		resD, err := mysqldb.QueryComments(s.ctx, req.VideoId, 10000, 0)
		if err != nil {
			return nil, err
		}
		res, err := pack.ChangeComments(s.ctx, resD)
		if err != nil {
			return nil, err
		}
		cacheReq := repository.NewRepositoryCache(1, req.VideoId).WithComments(resD)
		repository.ProducerCommentsCache(s.ctx, cacheReq)
		return res, nil
	})

	if err != nil {
		return nil, err
	}

	return v.([]*comment.Comment), nil
}
