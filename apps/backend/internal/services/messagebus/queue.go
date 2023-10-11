package messagebus

import (
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"
)

type queue struct {
	id          uuid.UUID
	subscribers map[uuid.UUID]subscriber
	messages    []message
	mu          sync.Mutex
}

func NewQueue() queue {
	q := queue{
		id:          uuid.New(),
		subscribers: make(map[uuid.UUID]subscriber),
		messages:    make([]message, 0),
	}

	return q
}

func (q *queue) Subscribe(subscriber subscriber) error {
	_, exists := q.subscribers[subscriber.GetID()]

	if exists {
		return fmt.Errorf("subscriber %v already registered", subscriber.GetID().String())
	}

	q.subscribers[subscriber.GetID()] = subscriber

	return nil
}

func (q *queue) GetID() uuid.UUID {
	return q.id
}

func (q *queue) HandleMessage(msg message) error {
	q.mu.Lock()
	q.messages = append(q.messages, msg)
	go q.publish()
	q.mu.Unlock()

	return nil
}

func (q *queue) publish() {
	q.mu.Lock()
	defer q.mu.Unlock()

	log.Printf("try to publish %v messages", len(q.messages))

	for i, msg := range q.messages {
		isPublished := false

		for !isPublished {

			for _, subscriber := range q.subscribers {
				if err := subscriber.HandleMessage(msg); err == nil {
					q.messages = append(q.messages[:i], q.messages[i+1:]...)
					isPublished = true
					break
				}
			}
		}

	}
}
