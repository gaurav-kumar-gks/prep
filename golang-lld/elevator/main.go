/*
Folder structure:

elevator/
├── main.go
├── models.go
├── errors.go
├── service.go
└── service_test.go

---
# main.go
/*
LLD Design for Elevator Control System typically tests:
1) OOP modelling of multiple entities and interactions
2) Dispatch algorithms (selection logic) and state management
3) Thread-safety for concurrent requests and movement simulation

Typical Requirements:
- Initialize building with F floors and N elevators
- CallElevator(floor, direction) returns assigned elevator ID
- SelectFloor(elevatorID, destination) adds destination to elevator stops
- Step() simulates one time-unit move for all elevators
- GetStatus() returns state (current floor, direction, stops) of each elevator

Typical Extensions & Solutions:
- Support emergency mode (redirect all elevators to ground floor)
- Dynamic addition/removal of elevators (scale-out/in)
- Prioritized scheduling (VIP vs normal floors)
- Partition floors into zones and assign elevator groups
- Expose APIs and persistence via DB or distributed coordination
*/

// ----- main.go -----
package main

import (
    "fmt"
    "time"
)

func main() {
    // Initialize service with 3 elevators, 10 floors
    svc := NewElevatorService(3, 10)

    // External calls
    eid, _ := svc.CallElevator(0, Up)     // call from ground to go up
    fmt.Println("Assigned Elevator:", eid)
    svc.SelectFloor(eid, 5)               // passenger selects floor 5

    // simulation loop
    for i := 0; i < 7; i++ {
        statuses := svc.GetStatus()
        fmt.Printf("Tick %d:\n", i)
        for _, st := range statuses {
            fmt.Println(st)
        }
        svc.Step()
        time.Sleep(200 * time.Millisecond)
    }

    // Another call
    eid2, _ := svc.CallElevator(7, Down)
    fmt.Println("Assigned Elevator:", eid2)
    svc.SelectFloor(eid2, 2)
    for i := 0; i < 10; i++ {
        svc.Step()
    }
    fmt.Println("Final statuses:", svc.GetStatus())
}
