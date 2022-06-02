package service

import (
	"context"
	"douyin/cmd/comment/dal/mysqldb"
	"douyin/cmd/comment/dal/redisdb"
	"douyin/cmd/comment/pack"
	"douyin/cmd/comment/repository"
	"douyin/kitex_gen/comment"
	"log"
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
		//TODO:log
		log.Fatal("redisdb err: ", err)
	}

	if status {
		return res.Comments, nil
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
		cacheReq := repository.NewRepositoryCache(1, req.VideoId).WithComments(res)
		repository.ProducerCommentsCache(s.ctx, cacheReq)
		return res, nil
	})

	if err != nil {
		return nil, err
	}

	return v.([]*comment.Comment), nil
}
