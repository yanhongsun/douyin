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
	status, err := redisdb.CheckCommentsCache(req.VideoId)

	if err != nil {
		//TODO:log
		log.Fatal("NewConsumer err: ", err)
	}

	if status {
		res, err := redisdb.GetCommentsCache(req.VideoId)
		if err == nil {
			return res, nil
		}
		log.Fatal("redis err: ", err)
	}

	key := strconv.FormatInt(req.VideoId, 10)

	v, err, _ := gQueryComments.Do(key, func() (interface{}, error) {
		resD, err := mysqldb.QueryComments(s.ctx, req.VideoId, 10000, 0)
		if err != nil {
			return nil, err
		}
		resD = pack.ReverseComments(resD)
		res := pack.ChangeComments(resD)
		repository.ProducerCommentsCache(1, req.VideoId, res, nil, -10001)
		return res, nil
	})

	if err != nil {
		return nil, err
	}

	return v.([]*comment.Comment), nil
}
