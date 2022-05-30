package repository

import (
	"context"
	"douyin/cmd/comment/dal/redisdb"
	"douyin/cmd/comment/pack/configdata"
	"douyin/kitex_gen/comment"
	"douyin/pkg/tracer"
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
	"github.com/cloudwego/kitex/pkg/klog"
)

type repositoryCache struct {
	// Type: 1-create_comments_cache,2-update_add,3-update_delete,4-create_comment_number_cache
	Type          int64              `json:"type"`
	VideoId       int64              `json:"video_id"`
	Comments      []*comment.Comment `json:"comments"`
	Comment       *comment.Comment   `json:"comment"`
	CommentId     int64              `json:"comment_id"`
	CommentNumber int64              `json:"comment_number"`
}

func ProducerCommentsCache(ctx context.Context, types, videoId int64, comments []*comment.Comment, comment *comment.Comment, commentId, commentNumber int64) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = configdata.KafkaConfig.TopicCommentNumber
	dataRepository := repositoryCache{
		Type:          types,
		VideoId:       videoId,
		Comments:      comments,
		Comment:       comment,
		CommentId:     commentId,
		CommentNumber: commentNumber,
	}
	data, err := json.Marshal(dataRepository)
	if err != nil {
		klog.Fatal(err)
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

	// context.WithValue()
}

func ConsumeCommentsCache(ctx context.Context) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{configdata.KafkaConfig.Host}, config)
	if err != nil {
		// TODO: log
		log.Fatal("NewConsumer err: ", err)
		return
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(configdata.KafkaConfig.TopicCommentNumber, 0, sarama.OffsetNewest)
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
			redisdb.AddCommentsCache(ctx, data.VideoId, data.Comments)
			continue
		} else if data.Type == 2 {
			redisdb.UpdateCommentsCache(ctx, data.VideoId, data.Comment)
			continue
		} else if data.Type == 3 {
			redisdb.DeleteCommentsCache(ctx, data.VideoId, data.CommentId)
			continue
		} else if data.Type == 4 {
			redisdb.AddCommentNumberCache(ctx, data.VideoId, data.CommentNumber)
			continue
		}
		log.Fatal("type is wrong")
	}
}
