package main

import (
    "sync"
    "time"
)

// SendFunc defines how to send a notification
type SendFunc func(n *Notification) error

// Dispatcher handles asynchronous notification delivery
type Dispatcher struct {
    sendFunc    SendFunc
    maxRetries  int
    backoff     time.Duration
    workers     int

    queue       chan *Notification
    deadLetter  []*Notification
    mu          sync.Mutex
    wg          sync.WaitGroup
    stopCh      chan struct{}
}

// NewDispatcher constructs a dispatcher
func NewDispatcher(workers, maxRetries int, backoff time.Duration, sf SendFunc) *Dispatcher {
    return &Dispatcher{
        sendFunc:   sf,
        maxRetries: maxRetries,
        backoff:    backoff,
        workers:    workers,
        queue:      make(chan *Notification, 100),
        stopCh:     make(chan struct{}),
    }
}

// Start launches worker goroutines
func (d *Dispatcher) Start() {
    for i := 0; i < d.workers; i++ {
        d.wg.Add(1)
        go d.worker()
    }
}

// Stop signals workers to finish and waits for completion
func (d *Dispatcher) Stop() {
    close(d.stopCh)
    d.wg.Wait()
}

// Enqueue adds a notification to the dispatch queue
func (d *Dispatcher) Enqueue(n *Notification) {
    select {
    case d.queue <- n:
    case <-d.stopCh:
    }
}

// DeadLetter returns failed notifications
func (d *Dispatcher) DeadLetter() []*Notification {
    d.mu.Lock()
    defer d.mu.Unlock()
    return append([]*Notification(nil), d.deadLetter...)
}

func (d *Dispatcher) worker() {
    defer d.wg.Done()
    for {
        select {
        case n := <-d.queue:
            d.process(n)
        case <-d.stopCh:
            return
        }
    }
}

func (d *Dispatcher) process(n *Notification) {
    for {
        err := d.sendFunc(n)
        if err == nil {
            // success
            return
        }
        n.Attempts++
        if n.Attempts > d.maxRetries {
            d.mu.Lock()
            d.deadLetter = append(d.deadLetter, n)
            d.mu.Unlock()
            return
        }
        time.Sleep(d.backoff)
    }
}
