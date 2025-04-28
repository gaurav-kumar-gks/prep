package main

import (
	"container/heap"
	"fmt"
)

type Direction string

const (
    Up   Direction = "UP"
    Down Direction = "DOWN"
    Idle Direction = "IDLE"
)

// FloorRequest represents a request to go to a specific floor
type FloorRequest struct {
    Floor     int
    Direction Direction
    index     int // used by heap.Interface
}

// PriorityQueue implements heap.Interface
type PriorityQueue struct {
    items []*FloorRequest
    ascending bool // true for min-heap (ascending), false for max-heap (descending)
}

// NewPriorityQueue creates a new priority queue with the specified ordering
func NewPriorityQueue(ascending bool) *PriorityQueue {
    pq := &PriorityQueue{
        items: make([]*FloorRequest, 0),
        ascending: ascending,
    }
    heap.Init(pq)
    return pq
}

func (pq PriorityQueue) Len() int { return len(pq.items) }

func (pq PriorityQueue) Less(i, j int) bool {
    if pq.ascending {
        return pq.items[i].Floor < pq.items[j].Floor
    }
    return pq.items[i].Floor > pq.items[j].Floor
}

func (pq PriorityQueue) Swap(i, j int) {
    pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
    pq.items[i].index = i
    pq.items[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
    n := len(pq.items)
    request := x.(*FloorRequest)
    request.index = n
    pq.items = append(pq.items, request)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := pq.items
    n := len(old)
    request := old[n-1]
    old[n-1] = nil // avoid memory leak
    request.index = -1
    pq.items = old[0 : n-1]
    return request
}

// Elevator represents an elevator car
type Elevator struct {
    ID           int
    CurrentFloor int
    Direction    Direction
    UpStops      *PriorityQueue   // stops when going up (ascending order)
    DownStops    *PriorityQueue   // stops when going down (descending order)
}

func NewElevator(id int) *Elevator {
    e := &Elevator{
        ID:           id,
        CurrentFloor: 0,
        Direction:    Idle,
        UpStops:      NewPriorityQueue(true),  // ascending order for up
        DownStops:    NewPriorityQueue(false), // descending order for down
    }
    return e
}

func (e Elevator) String() string {
    return fmt.Sprintf("Elevator %d: Floor=%d, Dir=%s, UpStops=%v, DownStops=%v", 
        e.ID, e.CurrentFloor, e.Direction, e.UpStops.items, e.DownStops.items)
}

// ElevatorStatus is used for reporting
type ElevatorStatus = Elevator