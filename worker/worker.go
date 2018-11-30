package worker

import (
	"encoding/json"
	log "github.com/golang/glog"
	"time"

	"github.com/whatvn/dqueue/helper"
	"github.com/whatvn/dqueue/queue"
	"github.com/whatvn/dqueue/models"
)

type Worker struct {
	queue queue.Queue
}

func NewWorker(queueType string) *Worker {
	log.Infof("worker type: %s started", queueType)
	w := &Worker{}

	switch queueType {
	case "NATS":
		w.queue = queue.NewNatsQueue()
		break
	case "STAN":
		w.queue = queue.NewStanQueue()
		break
	case "KAFKA":
		w.queue = queue.NewKafkaQueue()
		break
	default:
		panic(message.NotImplementError)
	}
	return w
}

func (w *Worker) Run() {
	for {
		now := helper.Now()
		msgList, err := message.GetMessagesByTimeStamp(now)
		if err != nil {
			log.Infof("cannot get message from database, error: ", err)
			time.Sleep(2 * time.Second)
			continue
		}

		for _, msg := range msgList {
			data, _ := json.Marshal(msg)
			log.Info("publishing message: ", msg)

			err = w.queue.PublishMessage(data)
			if err == nil {
				message.DeleteMessage(msg)
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}
