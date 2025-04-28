package main

import (
    "sync"
    "testing"
    "time"
)

func TestTokenBucketBasic(t *testing.T) {
    tb := NewTokenBucket(3, 1.0) // 3 tokens, refill 1/sec
    // consume all tokens
    for i := 0; i < 3; i++ {
        if !tb.Allow() {
            t.Errorf("expected allow at iteration %d", i)
        }
    }
    // next should be denied
    if tb.Allow() {
        t.Error("expected deny after tokens exhausted")
    }
    // wait 2 seconds -> 2 tokens
    time.Sleep(2 * time.Second)
    if !tb.Allow() || !tb.Allow() {
        t.Error("expected two allows after refill")
    }
}

func TestClientLimiterConcurrency(t *testing.T) {
    cl := NewClientLimiter(1, 10) // burst 1, high refill
    key := "k"
    wg := sync.WaitGroup{}
    success := 0
    mu := sync.Mutex{}
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            if cl.AllowClient(key) {
                mu.Lock()
                success++
                mu.Unlock()
            }
        }()
    }
    wg.Wait()
    if success != 1 {
        t.Errorf("expected exactly one success, got %d", success)
    }
}
