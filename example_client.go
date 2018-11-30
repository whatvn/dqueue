package main

import (
	"context"
	"log"
	"time"

	"github.com/whatvn/dqueue/protobuf"
	"github.com/whatvn/dqueue/wrapper"
)

func main() {
	client := wrapper.NewDelayQueueClient()
	message := &delayQueue.QueueRequest{
		Messsage:   "hello world",
		RetryCount: 3,
		Delay:      1,
	}
	ctx := context.Background()
	for {
		resp, err := client.Publish(ctx, message)
		time.Sleep(1 * time.Second)
		log.Println(resp, err)
	}
}
