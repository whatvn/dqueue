package queue

import (
	"fmt"

	log "github.com/golang/glog"
	"github.com/nats-io/go-nats-streaming"
	"github.com/whatvn/dqueue/helper"
)

type stanQueue struct {
	client stan.Conn
}

type stanConfig struct {
	Address, Port, ClusterID, ClientID string
}

func NewStanQueue() *stanQueue {
	var conf stanConfig
	helper.Config(&conf, "hosts", "stan")
	server := fmt.Sprintf("%s:%s", conf.Address, conf.Port)
	natsUrl := "nats://" + server
	log.Info("connecting to nats streaming : ", natsUrl)

	ns, err := stan.Connect(conf.ClusterID, conf.ClientID, stan.NatsURL(natsUrl))
	if err != nil {
		log.Error("cannot cannot to nats streaming , error: ", err)
		panic(err)
	}
	return &stanQueue{
		client: ns,
	}
}

func (queue *stanQueue) PublishMessage(message []byte) error {
	err := queue.client.Publish("pendingqueue", message)
	if err != nil {
		log.Error("cannot publish message", string(message), "error:", err)
		return err
	}
	return nil
}
