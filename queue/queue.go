package queue

type Queue interface {
	PublishMessage([]byte) error
}
