package main

import (
    "errors"
    "sync"
    "testing"
    "time"
)

func TestDispatcherRetriesAndDeadLetter(t *testing.T) {
    var mu sync.Mutex
    sent := make(map[string]int)
    // sendFunc tracks attempts and fails on first 3
    sendFunc := func(n *Notification) error {
        mu.Lock()
        sent[n.ID]++
        cnt := sent[n.ID]
        mu.Unlock()
        if cnt <= 3 {
            return errors.New("fail")
        }
        return nil
    }

    disp := NewDispatcher(2, 2, 100*time.Millisecond, sendFunc)
    disp.Start()
    n := &Notification{ID: "X", Recipient: "a@b", Message: "M"}
    disp.Enqueue(n)

    time.Sleep(1 * time.Second)
    disp.Stop()
    dl := disp.DeadLetter()
    if len(dl) != 1 || dl[0].ID != "X" {
        t.Errorf("expected X in dead-letter, got %v", dl)
    }
}

func TestDispatcherSuccess(t *testing.T) {
    sendFunc := func(n *Notification) error { return nil }
    disp := NewDispatcher(1, 1, 10*time.Millisecond, sendFunc)
    disp.Start()
    n := &Notification{ID: "Y", Recipient: "c@d", Message: "Hello"}
    disp.Enqueue(n)
    time.Sleep(50 * time.Millisecond)
    disp.Stop()
    dl := disp.DeadLetter()
    if len(dl) != 0 {
        t.Errorf("expected no dead-letter, got %v", dl)
    }
}
