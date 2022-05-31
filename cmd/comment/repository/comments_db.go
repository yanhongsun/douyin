package repository

import (
	"context"
	"douyin/cmd/comment/dal/mysqldb"
	"douyin/cmd/comment/pack"
	"douyin/cmd/comment/pack/configdata"
	"douyin/cmd/comment/rpc"
	"douyin/pkg/tracer"
	"encoding/json"
	"log"
	"strconv"

	"github.com/Shopify/sarama"
)

type repositoryCom struct {
	// Type:1-create,2-delete
	Type      int64            `json:"type"`
	Comment   *mysqldb.Comment `json:"comment"`
	CommentId int64            `json:"comment_id"`

	// can choose
	VideoId int64         `json:"video_id"`
	User    *rpc.UserInfo `json:"user"`
	UserId  int64         `json:user_id`
}

func ProducerComment(ctx context.Context, types int64, comment *mysqldb.Comment, commentId, videoId int64, user *rpc.UserInfo, userId int64) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = configdata.KafkaConfig.TopicComments

	dataRepository := repositoryCom{
		Type:      types,
		Comment:   comment,
		CommentId: commentId,
		VideoId:   videoId,
		User:      user,
		UserId:    userId,
	}

	data, err := json.Marshal(dataRepository)

	if err != nil {
		return err
	}
	msg.Value = sarama.StringEncoder(string(data))

	clients, err := sarama.NewSyncProducer([]string{configdata.KafkaConfig.Host}, config)
	client := tracer.NewSyncProducerFromClient(clients)
	client = client.WithContext(ctx)

	if err != nil {
		return err
	}
	defer client.Close()

	if _, _, err := client.SendMessage(msg); err != nil {
		return err
	}

	return nil
}

func ConsumeComments(ctx context.Context) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{configdata.KafkaConfig.Host}, config)
	if err != nil {
		// TODO: log
		log.Fatal("NewConsumer err: ", err)
		return
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(configdata.KafkaConfig.TopicComments, 0, sarama.OffsetNewest)
	if err != nil {
		// TODO: log
		log.Fatal("ConsumePartition err: ", err)
		return
	}
	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		res := message.Value
		var data repositoryCom
		err := json.Unmarshal([]byte(res), &data)
		if err != nil {
			// TODO: log
			log.Fatal("Json Unmarshal err: ", err)
			continue
		}
		if data.Type == 1 {
			textModeration, err := strconv.ParseBool(configdata.TencentCloudConfig.TextModeration)
			if err != nil {
				//log
				continue
			}
			moderationRes := "Pass"
			if textModeration {
				moderationRes, err = pack.CommentModeration(data.Comment.Content)
				if err != nil {
					log.Fatal("API err: ", err)
				}
			}
			if moderationRes == "Review" {
				data.Comment.State = false
			} else if moderationRes == "Block" {
				continue
			}
			err = mysqldb.CreateComment(ctx, data.Comment)
			if err != nil {
				// log
				continue
			}
			ProducerCommentsCache(ctx, 2, data.Comment.VideoID, nil, pack.ChangeComment(data.Comment, data.User), -10001, -10001)
			continue
		} else if data.Type == 2 {
			err = mysqldb.DeleteComment(ctx, data.CommentId, data.VideoId, data.UserId)
			if err != nil {
				continue
			}
			ProducerCommentsCache(ctx, 3, data.VideoId, nil, nil, data.CommentId, -10001)
			continue
		}
		log.Fatal("type is wrong")
	}
}
