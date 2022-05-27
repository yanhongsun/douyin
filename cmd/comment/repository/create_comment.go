package repository

import (
	"context"
	"douyin/cmd/comment/dal/mysqldb"
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
)

func ProducerCreateComment(comment *mysqldb.Comment) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "create_comment"
	data, err := json.Marshal(comment)
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

func ConsumeCreateComment() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		// TODO: log
		log.Fatal("NewConsumer err: ", err)
		return
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("create_comment", 0, sarama.OffsetNewest)
	if err != nil {
		// TODO: log
		log.Fatal("ConsumePartition err: ", err)
		return
	}
	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		res := message.Value
		var comment mysqldb.Comment
		err := json.Unmarshal([]byte(res), &comment)
		if err != nil {
			// TODO: log
			log.Fatal("Json Unmarshal err: ", err)
			continue
		}
		mysqldb.CreateComment(context.Background(), &comment)
	}
}
