package repository

type TaskConsumer interface {
	Consume() error
}
