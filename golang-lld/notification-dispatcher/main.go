/*
LLD Design for Notification/Email Dispatcher typically tests:
1) Asynchronous job queue design with worker pool
2) Retry logic and dead-letter handling for failures
3) Thread-safe queue operations and graceful shutdown

Typical Requirements:
- Accept notifications (email/SMS) for dispatch
- Dispatch asynchronously via worker pool
- Retry on failure up to `maxRetries` with backoff
- Maintain a dead-letter queue for permanently failed notifications
- Provide statistics: total sent, failed, in-flight

Typical Extensions & Solutions:
- **Distributed Queue**: use Kafka/RabbitMQ with idempotent producers and consumers
- **Persistent Outbox Pattern**: store notifications in DB and publish reliably
- **Exponential Backoff & Jitter**: to avoid thundering herd
- **Delayed Scheduling**: support scheduling notifications at a future timestamp
- **Metrics & Monitoring**: integrate with Prometheus/Grafana for queue depth, error rates
- **Circuit Breaker**: between dispatcher and email provider to handle provider outages
*/

// models.go


// errors.go


// service.go

// main.go
package main

import (
    "errors"
    "fmt"
    "time"
)

func main() {
    // Simulated send function: fails first two attempts
    sendFunc := func(n *Notification) error {
        if n.Attempts < 2 {
            fmt.Printf("Sending %s attempt %d... fail\n", n.ID, n.Attempts+1)
            return errors.New("temporary error")
        }
        fmt.Printf("Sending %s attempt %d... success\n", n.ID, n.Attempts+1)
        return nil
    }

    dispatcher := NewDispatcher(3, 5, 500*time.Millisecond, sendFunc)
    dispatcher.Start()

    // Enqueue notifications
    for i := 1; i <= 5; i++ {
        n := &Notification{ID: fmt.Sprintf("N%d", i), Recipient: "user@example.com", Message: "Hello"}
        dispatcher.Enqueue(n)
    }

    // Allow some time for processing
    time.Sleep(5 * time.Second)
    dispatcher.Stop()

    // Check dead-letter
    dl := dispatcher.DeadLetter()
    fmt.Println("Dead-letter notifications:", dl)
}

// service_test.go
