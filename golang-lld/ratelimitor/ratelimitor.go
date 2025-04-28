package main

import (
    "sync"
    "time"
)

// TokenBucket implements a classic token bucket rate limiter
type TokenBucket struct {
    mu         sync.Mutex
    capacity   float64       // maximum tokens
    tokens     float64       // current tokens
    refillRate float64       // tokens per second
    lastTime   time.Time     // last timestamp tokens were refilled
}

// NewTokenBucket constructs a token bucket with given capacity and refill rate
func NewTokenBucket(capacity int, refillRate float64) *TokenBucket {
    return &TokenBucket{
        capacity:   float64(capacity),
        tokens:     float64(capacity),
        refillRate: refillRate,
        lastTime:   time.Now(),
    }
}

// Allow checks if a token can be consumed now; returns true if allowed
func (tb *TokenBucket) Allow() bool {
    tb.mu.Lock()
    defer tb.mu.Unlock()
    now := time.Now()
    elapsed := now.Sub(tb.lastTime).Seconds()
    // refill tokens
    tb.tokens += elapsed * tb.refillRate
    if tb.tokens > tb.capacity {
        tb.tokens = tb.capacity
    }
    tb.lastTime = now
    if tb.tokens < 1.0 {
        return false
    }
    tb.tokens -= 1.0
    return true
}

// ClientLimiter manages multiple buckets keyed by client ID
type ClientLimiter struct {
    buckets   map[string]*TokenBucket
    mu        sync.Mutex
    capacity  int
    refillRate float64
}

// NewClientLimiter creates a manager for per-client buckets
func NewClientLimiter(capacity int, rate float64) *ClientLimiter {
    return &ClientLimiter{
        buckets:    make(map[string]*TokenBucket),
        capacity:   capacity,
        refillRate: rate,
    }
}

// AllowClient checks rate limit for given client key
func (cl *ClientLimiter) AllowClient(key string) bool {
    cl.mu.Lock()
    bucket, exists := cl.buckets[key]
    if !exists {
        bucket = NewTokenBucket(cl.capacity, cl.refillRate)
        cl.buckets[key] = bucket
    }
    cl.mu.Unlock()
    return bucket.Allow()
}
