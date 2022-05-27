package repository

import (
	"douyin/cmd/comment/dal/redisdb"
	"douyin/kitex_gen/comment"
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
)

type repositoryCache struct {
	VideoId int64              `json:"video_id"`
	Comment []*comment.Comment `json:"comment"`
}

func ProducerCreateCommentsCache(videoId int64, comment []*comment.Comment) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "create_comment_cache"
	dataRepository := repositoryCache{
		VideoId: videoId,
		Comment: comment,
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

func ConsumeCreateCommentsCache() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		// TODO: log
		log.Fatal("NewConsumer err: ", err)
		return
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("create_comment_cache", 0, sarama.OffsetNewest)
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
		redisdb.AddCommentsCache(data.VideoId, data.Comment)
	}
}
