package message

import (
	"github.com/whatvn/dqueue/protobuf"
	"github.com/whatvn/dqueue/helper"
)

type Message interface {
}

func NewMessage(queueRequest *delayQueue.QueueRequest) Message {
	switch helper.GetDbType() {
	case MYSQL:
		return &MySQLMessage{
			TimeStamp:  helper.NowPlus(queueRequest.Delay*queueRequest.RetryCount),
			Data:       queueRequest.Messsage,
			RetryCount: int(queueRequest.RetryCount),
			Delay:      int(queueRequest.Delay),
		}
	default:
		panic(NotImplementError)

	}
}