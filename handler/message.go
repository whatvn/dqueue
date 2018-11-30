package handler

import (
	"strconv"

	log "github.com/golang/glog"
	"github.com/whatvn/dqueue/models"
	"github.com/whatvn/dqueue/protobuf"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct{}


func (md *MessageHandler) GetAllMessages(c *gin.Context) {
	log.Info("Received Get all messages request")
	var (
		response = &delayQueue.QueryListMessagesResp{}
	)

	msgList, err := message.GetAllMessages()
	if err != nil {
		log.Error("cannot get all message, error: ", err)
		response.ReturnCode = message.Fail
		c.JSON(200, response)
	}

	log.Info("message list", msgList)

	for _, msg := range msgList {
		response.MsgList = append(response.MsgList, &delayQueue.MessageData{
			TimeStamp:  msg.TimeStamp,
			Delay:      int32(msg.Delay),
			RetryCount: int32(msg.RetryCount),
			Data:       msg.Data,
		})
	}

	response.ReturnCode = message.Success
	response.Message = message.ErrorMessage(message.Success)


	log.Info("response", response.ReturnCode)

	c.JSON(200, response)
}

func (md *MessageHandler) GetListMessage(c *gin.Context) {
	log.Info("Received Get list message request")

	offset, _ := strconv.Atoi(c.Param("offset"))
	limit, _ := strconv.Atoi(c.Param("limit"))

	var (
		response = &delayQueue.QueryListMessagesResp{}
	)

	msgList, err := message.GetMessageList(offset, limit)
	if err != nil {
		log.Error("cannot get message list, error: ", err)
		response.ReturnCode = message.Fail
		c.JSON(200, response)
	}

	for _, msg := range msgList {
		response.MsgList = append(response.MsgList, &delayQueue.MessageData{
			TimeStamp:  msg.TimeStamp,
			Delay:      int32(msg.Delay),
			RetryCount: int32(msg.RetryCount),
			Data:       msg.Data,
		})
	}

	response.ReturnCode = message.Success

	c.JSON(200, response)
}

func (md *MessageHandler) GetListMessageByData(c *gin.Context) {
	log.Info("Received Get list message by data request")
	data := c.Param("data")

	log.Info("data", data)

	var (
		response = &delayQueue.QueryListMessagesResp{}
	)

	msgList, err := message.GetMessageListByData(data)
	if err != nil {
		log.Error("cannot get message by pattern: ", data, "error: ", err)
		response.ReturnCode = message.Fail
		c.JSON(200, response)
		return
	}

	for _, msg := range msgList {
		response.MsgList = append(response.MsgList, &delayQueue.MessageData{
			TimeStamp:  msg.TimeStamp,
			Delay:      int32(msg.Delay),
			RetryCount: int32(msg.RetryCount),
			Data:       msg.Data,
		})
	}

	response.ReturnCode = message.Success

	c.JSON(200, response)
}

func (md *MessageHandler) ForceMessage(c *gin.Context) {
	log.Info("Received Get force message request")
	data := c.Param("data")

	var (
		response = &delayQueue.ReturnCommon{}
	)

	err := message.UpdateTimestampMessage(data)
	if err != nil {
		log.Error("cannot update message timestamp, error: ", err)
		response.ReturnCode = message.Fail
		c.JSON(200, response)
	}

	response.ReturnCode = message.Success

	c.JSON(200, response)
}
