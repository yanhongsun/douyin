package repository

import (
	"context"
	"douyin/cmd/comment/dal/mysqldb"
	"douyin/cmd/comment/dal/redisdb"
	"douyin/cmd/comment/pack/configdata"
	"douyin/cmd/comment/pack/zapcomment"
	"douyin/pkg/tracer"
	"encoding/json"

	"github.com/Shopify/sarama"
)

type repositoryCache struct {
	// Type: 1-create_comments_cache,2-update_add,3-update_delete,4-create_comment_number_cache
	Type    int64 `json:"type"`
	VideoId int64 `json:"video_id"`

	// can choose
	Comments      []*mysqldb.Comment `json:"comments,omitempty"`
	Comment       *mysqldb.Comment   `json:"comment,omitempty"`
	CommentId     int64              `json:"comment_id,omitempty"`
	CommentNumber int64              `json:"comment_number,omitempty"`
}

func NewRepositoryCache(types int64, videoId int64) *repositoryCache {
	return &repositoryCache{Type: types, VideoId: videoId}
}

func (c *repositoryCache) WithComments(comments []*mysqldb.Comment) *repositoryCache {
	c.Comments = comments
	return c
}

func (c *repositoryCache) WithComment(comment *mysqldb.Comment) *repositoryCache {
	c.Comment = comment
	return c
}

func (c *repositoryCache) WithCommentId(commentId int64) *repositoryCache {
	c.CommentId = commentId
	return c
}

func (c *repositoryCache) WithCommentNumber(commentNumber int64) *repositoryCache {
	c.CommentNumber = commentNumber
	return c
}

func ProducerCommentsCache(ctx context.Context, rerepositoryCache *repositoryCache) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = configdata.KafkaConfig.TopicCommentNumber
	data, err := json.Marshal(rerepositoryCache)
	if err != nil {
		zapcomment.Logger.Error("json Unmarshal err: " + err.Error())
		return err
	}
	msg.Value = sarama.StringEncoder(string(data))

	clients, err := sarama.NewSyncProducer([]string{configdata.KafkaConfig.Host}, config)
	client := tracer.NewSyncProducerFromClient(clients)
	client = client.WithContext(ctx)
	if err != nil {
		zapcomment.Logger.Error("kafka client err: " + err.Error())
		return err
	}
	defer client.Close()

	if _, _, err := client.SendMessage(msg); err != nil {
		zapcomment.Logger.Error("kafka send redis message err: " + err.Error())
		return err
	}

	zapcomment.Logger.Info("kafka send redis message succeeded")

	return nil
}

func ConsumeCommentsCache(ctx context.Context) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{configdata.KafkaConfig.Host}, config)
	if err != nil {
		zapcomment.Logger.Panic("NewConsumer err: " + err.Error())
		return
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(configdata.KafkaConfig.TopicCommentNumber, 0, sarama.OffsetNewest)
	if err != nil {
		zapcomment.Logger.Panic("ConsumePartition err: " + err.Error())
		return
	}
	defer partitionConsumer.Close()

	zapcomment.Logger.Info("kafka in redis initialization succeeded")

	for message := range partitionConsumer.Messages() {
		res := message.Value
		var data repositoryCache
		err := json.Unmarshal([]byte(res), &data)
		if err != nil {
			zapcomment.Logger.Error("Json Unmarshal err: " + err.Error())
			continue
		}
		if data.Type == 1 {
			err := redisdb.AddCommentsCache(ctx, data.VideoId, data.Comments)
			if err == nil {
				zapcomment.Logger.Info("redis add cache succeeded")
			}
			continue
		} else if data.Type == 2 {
			redisdb.UpdateCommentsCache(ctx, data.VideoId, data.Comment)
			if err == nil {
				zapcomment.Logger.Info("redis update cache succeeded")
			}
			continue
		} else if data.Type == 3 {
			redisdb.DeleteCommentsCache(ctx, data.VideoId, data.CommentId)
			if err == nil {
				zapcomment.Logger.Info("redis delete some cache succeeded")
			}
			continue
		} else if data.Type == 4 {
			redisdb.AddCommentNumberCache(ctx, data.VideoId, data.CommentNumber)
			if err == nil {
				zapcomment.Logger.Info("redis add number cache succeeded")
			}
			continue
		}
		zapcomment.Logger.Error("cache type is wrong")
	}
}
