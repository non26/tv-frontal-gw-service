package enqueue

import (
	"tvfrontalgwservice/app/processor"
)

type IEnqueue interface {
	Enqueue(path string, body []byte) error
	Close()
	Do() error
}

type Enqueue struct {
	QueueSize int
	JobQueue  chan [2]string
	processor processor.IProcessor
}

func NewEnqueue(queueSize int, processor processor.IProcessor) IEnqueue {
	jq := make(chan [2]string, queueSize)
	return &Enqueue{
		QueueSize: queueSize,
		JobQueue:  jq,
		processor: processor,
	}
}

func (e *Enqueue) Enqueue(path string, body []byte) error {
	e.JobQueue <- [2]string{path, string(body)}
	return nil
}

func (e *Enqueue) Do() error {
	for job := range e.JobQueue {
		e.processor.ProcessTask(job[0], []byte(job[1]))
	}
	return nil
}

func (e *Enqueue) Close() {
	close(e.JobQueue)
}
