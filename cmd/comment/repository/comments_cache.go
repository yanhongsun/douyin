package repository

import (
	"douyin/cmd/comment/dal/redisdb"
	"douyin/kitex_gen/comment"
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
	"github.com/cloudwego/kitex/pkg/klog"
)

type repositoryCache struct {
	// Type: 1-create,2-update_add,3-update_delete
	Type      int64              `json:"type"`
	VideoId   int64              `json:"video_id"`
	Comments  []*comment.Comment `json:"comments"`
	Comment   *comment.Comment   `json:"comment"`
	CommentId int64              `json:"comment_id"`
}

func ProducerCommentsCache(types, videoId int64, comments []*comment.Comment, comment *comment.Comment, commentId int64) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "comments_cache"
	dataRepository := repositoryCache{
		Type:      types,
		VideoId:   videoId,
		Comments:  comments,
		Comment:   comment,
		CommentId: commentId,
	}
	data, err := json.Marshal(dataRepository)
	if err != nil {
		klog.Fatal(err)
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

func ConsumeCommentsCache() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		// TODO: log
		log.Fatal("NewConsumer err: ", err)
		return
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("comments_cache", 0, sarama.OffsetNewest)
	if err != nil {
		// TODO: log
		log.Fatal("ConsumePartition err: ", err)
		return
	}
	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		res := message.Value
		var data repositoryCache
		err := json.Unmarshal([]byte(res), &data)
		if err != nil {
			// TODO: log
			log.Fatal("Json Unmarshal err: ", err)
			continue
		}
		if data.Type == 1 {
			redisdb.AddCommentsCache(data.VideoId, data.Comments)
			continue
		} else if data.Type == 2 {
			redisdb.UpdateCommentsCache(data.VideoId, data.Comment)
			continue
		} else if data.Type == 3 {
			redisdb.DeleteCommentsCache(data.VideoId, data.CommentId)
			continue
		}
		log.Fatal("type is wrong")
	}
}
