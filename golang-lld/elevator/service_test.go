package main

import (
    "testing"
)

func TestCallAndStep(t *testing.T) {
    svc := NewElevatorService(2, 5)
    eid, _ := svc.CallElevator(3, Up)
    if eid != 1 && eid != 2 { t.Errorf("invalid elevator assigned %d", eid) }
    // step until reaches floor 3
    for i:=0; i<4; i++ { svc.Step() }
    status := svc.GetStatus()[eid-1]
    if status.CurrentFloor != 3 { t.Errorf("expected floor 3, got %d", status.CurrentFloor) }
}

func TestMultipleStops(t *testing.T) {
    svc := NewElevatorService(1, 10)
    svc.CallElevator(5, Up)
    svc.SelectFloor(1, 2)
    // stops sequence: [5,2]
    floors := []int{1,2,3,4,5,4,3,2}
    for _, _ = range floors {
        svc.Step()
    }
    status := svc.GetStatus()[0]
    if status.CurrentFloor != 2 || status.Direction != Idle { t.Errorf("expected idle at 2, got %v", status) }
}

func TestInvalidElevator(t *testing.T) {
    svc := NewElevatorService(1, 5)
    if err := svc.SelectFloor(99, 3); err == nil { t.Error("expected error for invalid elevator") }
}
