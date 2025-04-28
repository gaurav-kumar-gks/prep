/*
LLD Design for Rate Limiter typically tests:
1) Algorithmic design for controlling request rates under high concurrency
2) Data structures and time-based calculations (token bucket, sliding window)
3) Thread-safety and performance considerations for hot-path operations

Typical Requirements:
- Allow up to `capacity` requests in burst
- Refill tokens at a constant rate (`refillRate` tokens per second)
- Provide an `Allow()` method returning immediately whether a request is permitted
- Support multiple clients/keys (e.g., per-user rate limits)

Typical Extensions & Solutions:
- **Distributed Rate Limiting**: use centralized store (Redis) with Lua-based atomic token operations or Consistent Hash to shard buckets
- **Sliding Window Log**: maintain timestamp queue per client to enforce more accurate limits over rolling window
- **Fixed Window with Soft Limits**: combine fixed window counters with request weighting to smooth edge bursts
- **Multi-tier Limits**: apply both per-user and global service limits
- **IP-based & Endpoint-based**: different rules per API endpoint or IP prefix via dynamic configuration
*/

package main

import (
    "fmt"
    "time"
)

func main() {
    // Example: 5 tokens capacity, refilling 1 token per second
    rl := NewTokenBucket(5, 1.0)
    fmt.Println("TokenBucket demo (5 capacity, 1/sec refill):")
    for i := 0; i < 7; i++ {
        allowed := rl.Allow()
        fmt.Printf("Request %d allowed? %v\n", i+1, allowed)
    }
    fmt.Println("Sleeping 3 seconds to refill...")
    time.Sleep(3 * time.Second)
    for i := 0; i < 3; i++ {
        allowed := rl.Allow()
        fmt.Printf("After refill request %d allowed? %v\n", i+1, allowed)
    }

    // ClientLimiter demo
    cl := NewClientLimiter(2, 0.5) // 2 tokens, 0.5/sec ref
    key := "user1"
    fmt.Println("ClientLimiter demo for user1 (2 cap, 0.5/sec):")
    for i := 0; i < 4; i++ {
        fmt.Printf("Client request %d allowed? %v\n", i+1, cl.AllowClient(key))
    }
}
