package handler

import (
	"context"

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
	msg := message.NewMessage(request)
	_, err := msg.Save()
	if err != nil {
		log.Error("cannot add message to database, message: ", msg, "error: ", err)
		response.ReturnCode = message.Fail
		response.Message = message.ErrorMessage(message.Fail)
		return nil
	}

	response.ReturnCode = message.Success
	response.Message = message.ErrorMessage(message.Success)
	return nil
}
