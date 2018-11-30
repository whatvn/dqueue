package handler

import (
	"context"

	"github.com/whatvn/dqueue/helper"
	"github.com/whatvn/dqueue/models"
	"github.com/whatvn/dqueue/protobuf"
	log "github.com/golang/glog"
)

type MicroHandler struct {
}

// NewMicroServiceHandler
func NewMicroServiceHandler() delayQueue.DelayQueueHandler {
	return &MicroHandler{}
}

func (handler *MicroHandler) Publish(ctx context.Context, request *delayQueue.QueueRequest, response *delayQueue.QueueResponse) error {
	newMessage := &message.Message{
		TimeStamp:  helper.NowPlus(request.Delay*request.RetryCount),
		Data:       request.Messsage,
		RetryCount: int(request.RetryCount),
		Delay:      int(request.Delay),
	}

	_, err := message.AddMessage(newMessage)
	if err != nil {
		log.Error("cannot add message to database, message: ", newMessage, "error: ", err)
		response.ReturnCode = message.Fail
		response.Message = message.ErrorMessage(message.Fail)
		return nil
	}

	response.ReturnCode = message.Success
	response.Message = message.ErrorMessage(message.Success)
	return nil
}
