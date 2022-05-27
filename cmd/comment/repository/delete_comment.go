package repository

import (
	"context"
	"douyin/cmd/comment/dal/mysqldb"
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
)

type repositoryDelete struct {
	CommentId int64 `json:"comment_id"`
	VideoId   int64 `json:"vedio_id"`
	UserId    int64 `json:"user_id"`
}

func ProducerDeleteComment(commentId, videoId, userId int64) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	deleteComment := repositoryDelete{
		CommentId: commentId,
		VideoId:   videoId,
		UserId:    userId,
	}

	msg := &sarama.ProducerMessage{}
	msg.Topic = "delete_comment"
	data, err := json.Marshal(deleteComment)
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

func ConsumeDeleteComment() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		// TODO: log
		log.Fatal("NewConsumer err: ", err)
		return
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("delete_comment", 0, sarama.OffsetNewest)
	if err != nil {
		// TODO: log
		log.Fatal("ConsumePartition err: ", err)
		return
	}
	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		res := message.Value
		var deleteComment repositoryDelete
		err := json.Unmarshal([]byte(res), &deleteComment)
		if err != nil {
			// TODO: log
			log.Fatal("Json Unmarshal err: ", err)
			continue
		}
		mysqldb.DeleteComment(context.Background(), deleteComment.CommentId, deleteComment.VideoId, deleteComment.UserId)
	}
}
