package repository

import (
	"context"
	"douyin/cmd/comment/dal/mysqldb"
	"douyin/cmd/comment/pack"
	"douyin/cmd/comment/pack/configdata"
	"douyin/cmd/comment/pack/zapcomment"
	"douyin/cmd/comment/rpc"
	"douyin/pkg/tracer"
	"encoding/json"
	"strconv"

	"github.com/Shopify/sarama"
)

type repositoryCom struct {
	// Type:1-create,2-delete
	Type int64 `json:"type"`

	// can choose
	Comment   *mysqldb.Comment `json:"comment,omitempty"`
	User      *rpc.UserInfo    `json:"user,omitempty"`
	UserId    int64            `json:"user_id,omitempty"`
	VideoId   int64            `json:"video_id,omitempty"`
	CommentId int64            `json:"comment_id,omitempty"`
}

func NewRepositoryCom(types int64) *repositoryCom {
	return &repositoryCom{Type: types}
}

func (db *repositoryCom) WithComment(comment *mysqldb.Comment) *repositoryCom {
	db.Comment = comment
	return db
}

func (db *repositoryCom) WithUser(user *rpc.UserInfo) *repositoryCom {
	db.User = user
	return db
}

func (db *repositoryCom) WithVideoId(videoId int64) *repositoryCom {
	db.VideoId = videoId
	return db
}

func (db *repositoryCom) WithUserId(userId int64) *repositoryCom {
	db.UserId = userId
	return db
}

func (db *repositoryCom) WithCommentId(commentId int64) *repositoryCom {
	db.CommentId = commentId
	return db
}

func ProducerComment(ctx context.Context, repositoryCom *repositoryCom) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = configdata.KafkaConfig.TopicComments

	data, err := json.Marshal(repositoryCom)

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
		zapcomment.Logger.Error("kafka send mysql message err: " + err.Error())
		return err
	}

	zapcomment.Logger.Error("kafka send mysql message succeeded")

	return nil
}

func ConsumeComments(ctx context.Context) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{configdata.KafkaConfig.Host}, config)
	if err != nil {
		zapcomment.Logger.Panic("NewConsumer err: " + err.Error())
		return
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(configdata.KafkaConfig.TopicComments, 0, sarama.OffsetNewest)
	if err != nil {
		zapcomment.Logger.Panic("ConsumePartition err: " + err.Error())
		return
	}
	defer partitionConsumer.Close()
	zapcomment.Logger.Info("kafka in mysql initialization succeeded")

	for message := range partitionConsumer.Messages() {
		res := message.Value
		var data repositoryCom
		err := json.Unmarshal([]byte(res), &data)
		if err != nil {
			zapcomment.Logger.Error("json Unmarshal err: " + err.Error())
			continue
		}
		if data.Type == 1 {
			textModeration, err := strconv.ParseBool(configdata.TencentCloudConfig.TextModeration)
			if err != nil {
				zapcomment.Logger.Error("strcov error in : " + err.Error() + " : string - " + configdata.TencentCloudConfig.TextModeration)
				continue
			}
			moderationRes := "Pass"
			if textModeration {
				moderationRes, err = pack.CommentModeration(data.Comment.Content)
				if err != nil {
					zapcomment.Logger.Error("Tencent API err" + err.Error())
					moderationRes = "Review"
				}
			}
			if moderationRes == "Review" {
				data.Comment.State = false
			} else if moderationRes == "Block" {
				continue
			}
			err = mysqldb.CreateComment(ctx, data.Comment)
			if err != nil {
				zapcomment.Logger.Error("mysql commentId " + strconv.Itoa(int(data.CommentId)) + " create err" + err.Error())
			}

			cacheReq := NewRepositoryCache(2, data.Comment.VideoID).WithComment(pack.ChangeComment(data.Comment, data.User))
			ProducerCommentsCache(ctx, cacheReq)
			if err == nil {
				zapcomment.Logger.Error("mysql commentId " + strconv.Itoa(int(data.CommentId)) + " create succeeded")
			}
			continue
		} else if data.Type == 2 {
			err = mysqldb.DeleteComment(ctx, data.CommentId, data.VideoId, data.UserId)
			if err != nil {
				zapcomment.Logger.Error("mysql commentId " + strconv.Itoa(int(data.CommentId)) + " delete err" + err.Error())
			}
			cacheReq := NewRepositoryCache(3, data.VideoId).WithCommentId(data.CommentId)
			err := ProducerCommentsCache(ctx, cacheReq)
			if err == nil {
				zapcomment.Logger.Error("mysql commentId " + strconv.Itoa(int(data.CommentId)) + " delete succeeded")
			}
			continue
		}
		zapcomment.Logger.Panic("mysql type is wrong")
	}
}
