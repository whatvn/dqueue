package main

import (
	"context"
	"log"
	"time"

	"github.com/whatvn/dqueue/protobuf"
	"github.com/whatvn/dqueue/wrapper"
	"flag"
	"github.com/spf13/cast"
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Set("v", "2")
	flag.Parse()
	client := wrapper.NewDelayQueueClient()

	i := 1
	for {
		message := &delayQueue.QueueRequest{
			Messsage:   "hello world" + cast.ToString(i),
			RetryCount: 3,
			Delay:      3,
		}
		ctx := context.Background()
		resp, err := client.Publish(ctx, message)
		i += 1
		time.Sleep(100 * time.Millisecond)
		log.Println(resp, err)
	}
}
