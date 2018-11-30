package queue

import "github.com/nats-io/go-nats-streaming"

type NatsStreaming struct {
	client stan.Conn
	queueName string
}

func (n *NatsStreaming) Publish(message []byte)  {
	n.client.Publish(n.queueName, message)
}