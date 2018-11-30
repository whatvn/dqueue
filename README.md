# dqueue

dqueue is a delay queue written in Golang. dqueue was born because the need of a simple queuing layer which support delay before processing message.
A message pushed into dqueue will be stay in dqueue (which stored in database - currently mysql) until delay time reached. Delay time is client defined, in second. Message has following format: 

```json
message MessageData {
    int32 delay = 1;
    int32 retryCount = 2;
    string data = 3;
}
```

For example:

```json
{
  "delay": 1,
  "data": "hello world",
  "retryCount": 1
}
```

dqueue will delay message for 1*1 second, then push to message queue backend defined in configuration file, directive `queueType`. Then if subscriber fail to process message, eq: fail to store to database, fail to call or update an API, subscriber can then push back message into dqueue with `retryCount = currentRetryCount + 1`, message will be delay 1*2 second before published to queue backend, and so on

```json
{
  "delay": 1,
  "data": "hello world",
  "retryCount": 2
}
```

# Architecture 

![](https://raw.githubusercontent.com/whatvn/dqueue/master/diagram/diagram.jpg)

dqueue is written in golang, using go-micro framework, you can start as much as possible instance to scale dqueue. Client can also use go-micro client to work with dqueue server (include in this project)

Currently dqueue uses mysql as intermediate layer, and support multiple message queue system: nats, nats streaming, kafka

One can extend dqueue by adding other message queue system by implement `queue` interface: 

```go
package queue

type Queue interface {
	PublishMessage([]byte) error
}
```

see nats for example. 

One can also extend dqueue to support other intermediate database by implement `database` interface: 

```go
type Database interface {
	Init() error
}
```

dqueue also has apis to support listing, monitoring, and forcing a message to be in queue immediately using `force` method


# usage 

```go

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

```

#license 

BSD License 