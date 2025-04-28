package main

import (
	"container/heap"
	"math"
	"sync"
)

// ElevatorService manages multiple elevators
type ElevatorService struct {
	mu         sync.Mutex
	elevators  map[int]*Elevator
	numFloors  int
}

// NewElevatorService initializes N elevators at floor 0
func NewElevatorService(nElevators, numFloors int) *ElevatorService {
	es := &ElevatorService{elevators: make(map[int]*Elevator), numFloors: numFloors}
	for i := 1; i <= nElevators; i++ {
		es.elevators[i] = NewElevator(i)
	}
	return es
}

// CallElevator assigns the best elevator to a pickup request
func (es *ElevatorService) CallElevator(floor int, dir Direction) (int, error) {
	es.mu.Lock()
	defer es.mu.Unlock()
	
	bestID, bestDist := -1, math.MaxInt32
	for id, e := range es.elevators {
		dist := es.estimateDistance(e, floor, dir)
		if dist < bestDist {
			bestDist, bestID = dist, id
		}
	}
	
	e := es.elevators[bestID]
	request := &FloorRequest{Floor: floor, Direction: dir}
	
	// Add to appropriate queue based on direction and current position
	if floor > e.CurrentFloor || (floor == e.CurrentFloor && dir == Up) {
		heap.Push(e.UpStops, request)
	} else {
		heap.Push(e.DownStops, request)
	}
	
	es.updateDirection(e)
	return bestID, nil
}

// SelectFloor adds an in-elevator floor request
func (es *ElevatorService) SelectFloor(elevatorID, floor int) error {
	es.mu.Lock()
	defer es.mu.Unlock()
	
	e, ok := es.elevators[elevatorID]
	if !ok {
		return ErrInvalidElevator
	}
	
	request := &FloorRequest{Floor: floor, Direction: e.Direction}
	if e.Direction == Up || (e.Direction == Idle && floor > e.CurrentFloor) {
		heap.Push(e.UpStops, request)
	} else {
		heap.Push(e.DownStops, request)
	}
	
	es.updateDirection(e)
	return nil
}

// Step moves each elevator one step towards its next stop
func (es *ElevatorService) Step() {
	es.mu.Lock()
	defer es.mu.Unlock()
	
	for _, e := range es.elevators {
		if e.UpStops.Len() == 0 && e.DownStops.Len() == 0 {
			e.Direction = Idle
			continue
		}
		
		var targetFloor int
		var hasTarget bool
		
		// Get next target based on current direction
		if e.Direction == Up {
			if e.UpStops.Len() > 0 {
				targetFloor = e.UpStops.items[0].Floor
				hasTarget = true
			} else if e.DownStops.Len() > 0 {
				// Switch to down direction if no more up stops
				e.Direction = Down
				targetFloor = e.DownStops.items[0].Floor
				hasTarget = true
			}
		} else if e.Direction == Down {
			if e.DownStops.Len() > 0 {
				targetFloor = e.DownStops.items[0].Floor
				hasTarget = true
			} else if e.UpStops.Len() > 0 {
				// Switch to up direction if no more down stops
				e.Direction = Up
				targetFloor = e.UpStops.items[0].Floor
				hasTarget = true
			}
		} else { // Idle
			if e.UpStops.Len() > 0 {
				e.Direction = Up
				targetFloor = e.UpStops.items[0].Floor
				hasTarget = true
			} else if e.DownStops.Len() > 0 {
				e.Direction = Down
				targetFloor = e.DownStops.items[0].Floor
				hasTarget = true
			}
		}
		
		if !hasTarget {
			e.Direction = Idle
			continue
		}
		
		// Move towards target
		if e.CurrentFloor < targetFloor {
			e.CurrentFloor++
		} else if e.CurrentFloor > targetFloor {
			e.CurrentFloor--
		}
		
		// Check if we've reached the target
		if e.CurrentFloor == targetFloor {
			if e.Direction == Up && e.UpStops.Len() > 0 {
				heap.Pop(e.UpStops)
			} else if e.Direction == Down && e.DownStops.Len() > 0 {
				heap.Pop(e.DownStops)
			}
			
			// Update direction based on remaining stops
			es.updateDirection(e)
		}
	}
}

// GetStatus returns a snapshot of all elevators
func (es *ElevatorService) GetStatus() []ElevatorStatus {
	es.mu.Lock()
	defer es.mu.Unlock()
	
	statuses := make([]ElevatorStatus, 0, len(es.elevators))
	for _, e := range es.elevators {
		statuses = append(statuses, *e)
	}
	return statuses
}

// estimateDistance heuristically scores an elevator for assignment
func (es *ElevatorService) estimateDistance(e *Elevator, floor int, dir Direction) int {
	if e.Direction == Idle {
		return abs(e.CurrentFloor - floor)
	}
	
	// If moving towards request in same direction
	if (e.Direction == Up && floor >= e.CurrentFloor) || 
	   (e.Direction == Down && floor <= e.CurrentFloor) {
		return abs(e.CurrentFloor - floor)
	}
	
	// Calculate distance including direction change
	return abs(e.CurrentFloor-floor) + es.numFloors
}

func (es *ElevatorService) updateDirection(e *Elevator) {
	if e.UpStops.Len() == 0 && e.DownStops.Len() == 0 {
		e.Direction = Idle
		return
	}
	
	if e.Direction == Up && e.UpStops.Len() == 0 {
		e.Direction = Down
	} else if e.Direction == Down && e.DownStops.Len() == 0 {
		e.Direction = Up
	} else if e.Direction == Idle {
		if e.UpStops.Len() > 0 {
			e.Direction = Up
		} else if e.DownStops.Len() > 0 {
			e.Direction = Down
		}
	}
}

func abs(x int) int { if x<0 { return -x }; return x }