package main

import (
    "sync"
    "time"
)

type entry struct {
    value  interface{}
    expiry time.Time
}

type KVStore struct {
    mu    sync.RWMutex
    store map[string]*entry
}

// NewKVStore initializes the store and starts cleanup goroutine
func NewKVStore() *KVStore {
    kv := &KVStore{store: make(map[string]*entry)}
    go kv.cleanupExpired()
    return kv
}

// Put stores a value without expiration
func (kv *KVStore) Put(key string, value interface{}) {
    kv.mu.Lock()
    defer kv.mu.Unlock()
    kv.store[key] = &entry{value: value}
}

// PutWithTTL stores a value with time-to-live
func (kv *KVStore) PutWithTTL(key string, value interface{}, ttl time.Duration) {
    kv.mu.Lock()
    defer kv.mu.Unlock()
    kv.store[key] = &entry{value: value, expiry: time.Now().Add(ttl)}
}

// Get retrieves a value or error if not found/expired
func (kv *KVStore) Get(key string) (interface{}, error) {
    kv.mu.RLock()
    e, ok := kv.store[key]
    kv.mu.RUnlock()
    if !ok {
        return nil, ErrKeyNotFound
    }
    if !e.expiry.IsZero() && time.Now().After(e.expiry) {
        kv.mu.Lock()
        delete(kv.store, key)
        kv.mu.Unlock()
        return nil, ErrKeyNotFound
    }
    return e.value, nil
}

// Delete removes a key or returns error if absent
func (kv *KVStore) Delete(key string) error {
    kv.mu.Lock()
    defer kv.mu.Unlock()
    if _, exists := kv.store[key]; !exists {
        return ErrKeyNotFound
    }
    delete(kv.store, key)
    return nil
}

// cleanupExpired periodically removes expired entries
func (kv *KVStore) cleanupExpired() {
    ticker := time.NewTicker(time.Second)
    for now := range ticker.C {
        kv.mu.Lock()
        for k, e := range kv.store {
            if !e.expiry.IsZero() && now.After(e.expiry) {
                delete(kv.store, k)
            }
        }
        kv.mu.Unlock()
    }
}