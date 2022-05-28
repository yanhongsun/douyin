package repository

import (
	"context"
	"douyin/cmd/comment/dal/mysqldb"
	"douyin/cmd/comment/pack"
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
)

type repositoryCom struct {
	// Type:1-create,2-delete
	Type      int64            `json:"type"`
	Comment   *mysqldb.Comment `json:"comment"`
	CommentId int64            `json:"comment_id"`
	VideoId   int64            `json:"video_id"`
	UserId    int64            `json:"user_id"`
}

func ProducerComment(types int64, comment *mysqldb.Comment, commentId, videoId, userId int64) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "comment"

	dataRepository := repositoryCom{
		Type:      types,
		Comment:   comment,
		CommentId: commentId,
		VideoId:   videoId,
		UserId:    userId,
	}

	data, err := json.Marshal(dataRepository)
	if err != nil {
		return err
	}
	msg.Value = sarama.StringEncoder(string(data))

	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		return err
	}
	defer client.Close()

	if _, _, err := client.SendMessage(msg); err != nil {
		return err
	}

	return nil
}

func ConsumeComment() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		// TODO: log
		log.Fatal("NewConsumer err: ", err)
		return
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("comment", 0, sarama.OffsetNewest)
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
			err = mysqldb.CreateComment(context.Background(), data.Comment)
			if err != nil {
				continue
			}
			ProducerCommentsCache(2, data.Comment.VideoID, nil, pack.ChangeComment(data.Comment), -10001)
			continue
		} else if data.Type == 2 {
			err = mysqldb.DeleteComment(context.Background(), data.CommentId, data.VideoId, data.UserId)
			if err != nil {
				continue
			}
			ProducerCommentsCache(3, data.VideoId, nil, nil, -10001)
			continue
		}
		log.Fatal("type is wrong")
	}
}
