package main

import (
    "fmt"
    "sync"
    "testing"
    "time"
)

func TestPutGetDelete(t *testing.T) {
    kv := NewKVStore()
    kv.Put("k1", "v1")
    v, err := kv.Get("k1")
    if err != nil || v != "v1" {
        t.Fatalf("expected v1, got %v, err %v", v, err)
    }
    if err := kv.Delete("k1"); err != nil {
        t.Fatalf("delete error: %v", err)
    }
    if _, err := kv.Get("k1"); err != ErrKeyNotFound {
        t.Fatalf("expected ErrKeyNotFound, got %v", err)
    }
}

func TestTTLExpiry(t *testing.T) {
    kv := NewKVStore()
    kv.PutWithTTL("t1", "v", 1*time.Second)
    v, err := kv.Get("t1")
    if err != nil || v != "v" {
        t.Fatalf("expected v, got %v, err %v", v, err)
    }
    time.Sleep(2 * time.Second)
    if _, err := kv.Get("t1"); err != ErrKeyNotFound {
        t.Fatalf("expected TTL expiry, got %v", err)
    }
}

func TestConcurrentAccess(t *testing.T) {
    kv := NewKVStore()
    var wg sync.WaitGroup
    // concurrent writes
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            kv.Put(fmt.Sprintf("key%d", i), i)
        }(i)
    }
    // concurrent reads
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            kv.Get(fmt.Sprintf("key%d", i))
        }(i)
    }
    wg.Wait()
}
