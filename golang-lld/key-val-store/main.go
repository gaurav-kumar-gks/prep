package main

import (
    "fmt"
    "time"
)

func main() {
    kv := NewKVStore()
    kv.Put("foo", "bar")
    v, err := kv.Get("foo")
    fmt.Println(v, err) // bar, <nil>

    kv.PutWithTTL("baz", 42, 2*time.Second)
    v, err = kv.Get("baz")
    fmt.Println(v, err) // 42, <nil>

    time.Sleep(3 * time.Second)
    v, err = kv.Get("baz")
    fmt.Println(v, err) // <nil>, key not found

    if err := kv.Delete("foo"); err != nil {
        fmt.Println(err)
    }
    if _, err := kv.Get("foo"); err != ErrKeyNotFound {
        fmt.Println("expected deletion")
    }
}
