package queue

import (

	"time"
	log "github.com/golang/glog"
	"github.com/Shopify/sarama"
	kafka "github.com/Shopify/sarama"
	"github.com/whatvn/dqueue/helper"
)

type kafkaConfig struct {
	Host, Port, Topic string
}
type kafkaQueue struct {
	client Writer
}

func (q *kafkaQueue) PublishMessage(message []byte) error {
	return q.client.WriteRaw(message)
}

func NewKafkaQueue() *kafkaQueue {
	var conf kafkaConfig
	helper.Config(&conf, "hosts", "kafka")
	writer := createWriter(conf.Host+":"+conf.Port, conf.Topic)
	return &kafkaQueue{
		client: writer,
	}
}

//Writer ...
type Writer interface {
	WriteRaw([]byte) error
	Write(kafka.Encoder)
}
type writer struct {
	topic    string
	producer kafka.SyncProducer
}

//CreateWriter ....
func createWriter(addr, topic string) Writer {
	cfg := kafka.NewConfig()
	cfg.Producer.RequiredAcks = kafka.WaitForLocal
	cfg.Producer.Flush.Frequency = 50 * time.Millisecond
	producer, err := kafka.NewSyncProducer([]string{addr}, cfg)
	if err != nil {
		log.Error("kafka connection error: ", err)
		panic(err)
	}

	return &writer{topic: topic,
		producer: producer}

}
func (w *writer) Write(v kafka.Encoder) {
	msg := &kafka.ProducerMessage{
		Topic: w.topic,
		Value: v,
	}
	w.producer.SendMessage(msg)
}
func (w *writer) WriteRaw(v []byte) error {
	msg := &kafka.ProducerMessage{
		Topic: w.topic,
		Value: sarama.ByteEncoder(v),
	}
	_, _, err := w.producer.SendMessage(msg)
	return err
}
