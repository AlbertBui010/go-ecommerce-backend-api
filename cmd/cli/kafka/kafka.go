package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

var (
	kafkaUrl   = "localhost:19092"
	kafkaTopic = "user_topic_vip"
)

// for producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{}, // load balacing
	}
}

// for consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers, // []string{"localhost:111", ""}
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.LastOffset, // read from last offset, -Use FirstOffset to read first offset.
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMsg

	return &s
}

// Producer
func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s.Message

	jsonBody, _ := json.Marshal(body)
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err":    "",
		"action": "Action successfully",
	})
}

// Consumer listening
func RegisterConsumer(id int) {
	// One consumer group
	// kafkaGroupId := "consumer-group"

	//
	kafkaGroupId := fmt.Sprintf("consumer-group-%d", id)

	reader := getKafkaReader(kafkaUrl, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf(":::Consumer (%d) listening:::\n", id)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer (%d) error:: %v", id, err)
		}
		fmt.Printf("Consumer (%d) listening topic:%v, partition:%v, offset:%v, time:%d %s = %s\n",
			id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaUrl, kafkaTopic)
	defer kafkaProducer.Close()

	r.POST("action/stock", actionStock)

	// Register 2 consumer (listening)
	go RegisterConsumer(1)
	go RegisterConsumer(2)
	go RegisterConsumer(3)

	r.Run(":8999")
}
