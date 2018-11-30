// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: service.proto

package delayQueue

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for DelayQueue service

type DelayQueueService interface {
	Publish(ctx context.Context, in *QueueRequest, opts ...client.CallOption) (*QueueResponse, error)
}

type delayQueueService struct {
	c    client.Client
	name string
}

func NewDelayQueueService(name string, c client.Client) DelayQueueService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "delayQueue"
	}
	return &delayQueueService{
		c:    c,
		name: name,
	}
}

func (c *delayQueueService) Publish(ctx context.Context, in *QueueRequest, opts ...client.CallOption) (*QueueResponse, error) {
	req := c.c.NewRequest(c.name, "DelayQueue.Publish", in)
	out := new(QueueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DelayQueue service

type DelayQueueHandler interface {
	Publish(context.Context, *QueueRequest, *QueueResponse) error
}

func RegisterDelayQueueHandler(s server.Server, hdlr DelayQueueHandler, opts ...server.HandlerOption) error {
	type delayQueue interface {
		Publish(ctx context.Context, in *QueueRequest, out *QueueResponse) error
	}
	type DelayQueue struct {
		delayQueue
	}
	h := &delayQueueHandler{hdlr}
	return s.Handle(s.NewHandler(&DelayQueue{h}, opts...))
}

type delayQueueHandler struct {
	DelayQueueHandler
}

func (h *delayQueueHandler) Publish(ctx context.Context, in *QueueRequest, out *QueueResponse) error {
	return h.DelayQueueHandler.Publish(ctx, in, out)
}