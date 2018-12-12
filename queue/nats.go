package queue

import (
	"fmt"
	"time"
	log "github.com/golang/glog"
	"github.com/nats-io/go-nats"
	"github.com/whatvn/dqueue/helper"
)

type natsQueue struct {
	client *nats.Conn
}

type natsConfig struct {
	Address, Port string
}

func NewNatsQueue() *natsQueue {
	var conf natsConfig
	helper.Config(&conf, "hosts", "nats")
	server := fmt.Sprintf("%s:%s", conf.Address, conf.Port)
	log.Info("nats config: ", conf)
	opts := nats.Options{
		Url: server,
		// Servers:        servers,
		AllowReconnect: true,
		MaxReconnect:   10,
		ReconnectWait:  5 * time.Second,
		Timeout:        1 * time.Second,
	}

	nc, err := opts.Connect()
	if err != nil {
		log.Error("nats connection error: ", err)
		panic(err)
	}

	return &natsQueue{
		client: nc,
	}
}

func (queue *natsQueue) PublishMessage(message []byte) error {
	err := queue.client.Publish("pendingqueue", message)
	log.Info("publish message: ", string(message))
	if err != nil {
		log.Error("cannot publish message", string(message), "error:", err)
		return err
	}
	return nil
}
