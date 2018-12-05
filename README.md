# dqueue

dqueue is a delay queue written in Golang. dqueue was born because the need of a simple queuing layer which support delay before processing message.
A message pushed into dqueue will be stay in dqueue (which stored in database - currently mysql) until delay time reached. Delay time is client defined, in second. Message has following format: 

```json
{
  "delay": integer,
  "data": string,
  "retryCount": integer
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

dqueue will delay message for `1 * 1 = 1 second`, then push to message queue backend defined in configuration file, directive `queueType`. Then if subscriber fail to process message, eq: fail to store to database, fail to call or update an API, subscriber can then push back message into dqueue with `retryCount = currentRetryCount + 1`, message will be delay `1 * 2 = 2 seconds` before published to queue backend, and so on

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

Install dqueue

```bash
go get github.com/whatvn/dqueue
cd $GOPATH/src/github.com/whatvn/dqueue
go build main.go
```

change configuration file according to your system: 

```json
{
  "queueType": "NATS",
  "dbType": "mysql",
  "hosts": {
    "nats": {
      "address": "0.0.0.0",
      "port": "4222"
    },
    "stan": {
      "address": "0.0.0.0",
      "port": "4222",
      "clusterID": "test-cluster",
      "clientID": "retry-worker"
    },
    "mysql": {
      "address": "127.0.0.1",
      "port": "3306",
      "user": "root",
      "password": "123456",
      "database": "delay_queue"
    }
  }
}
```
where:
    - **queueType** is queue server backend
    - **dbType** is intermediate database server, currenly only mysql is supported


To publish a message into dqueue
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

To subscribe a message from your queue backend, use queue backend client 
# License 

BSD License 