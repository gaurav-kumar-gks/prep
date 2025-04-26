package main

import (
	"fmt"
	"sync"
	"time"
)

// Message represents a message that can be published
type Message struct {
	Topic   string
	Content string
}

// PubSub implements a simple publish-subscribe system
type PubSub struct {
	subscribers map[string][]chan Message
	mu          sync.RWMutex
}

// NewPubSub creates a new PubSub instance
func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan Message),
	}
}

// Subscribe allows a client to subscribe to a topic
func (ps *PubSub) Subscribe(topic string) chan Message {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	ch := make(chan Message, 1)
	ps.subscribers[topic] = append(ps.subscribers[topic], ch)
	return ch
}

// Publish sends a message to all subscribers of a topic
func (ps *PubSub) Publish(msg Message) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	
	for _, ch := range ps.subscribers[msg.Topic] {
		go func(ch chan Message) {
			ch <- msg
		}(ch)
	}
}

func main() {
	// Create a new PubSub instance
	ps := NewPubSub()

	// Create subscribers
	sub1 := ps.Subscribe("news")
	sub2 := ps.Subscribe("news")
	sub3 := ps.Subscribe("sports")

	// Start a goroutine to publish messages
	go func() {
		topics := []string{"news", "sports"}
		for i := 0; i < 5; i++ {
			topic := topics[i%len(topics)]
			msg := Message{
				Topic:   topic,
				Content: fmt.Sprintf("Message %d for %s", i+1, topic),
			}
			ps.Publish(msg)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Collect messages from subscribers
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for msg := range sub1 {
			fmt.Printf("Subscriber 1 received: %s on topic %s\n", msg.Content, msg.Topic)
		}
	}()

	go func() {
		defer wg.Done()
		for msg := range sub2 {
			fmt.Printf("Subscriber 2 received: %s on topic %s\n", msg.Content, msg.Topic)
		}
	}()

	go func() {
		defer wg.Done()
		for msg := range sub3 {
			fmt.Printf("Subscriber 3 received: %s on topic %s\n", msg.Content, msg.Topic)
		}
	}()

	// Wait for a while to see the messages
	time.Sleep(1 * time.Second)
} 