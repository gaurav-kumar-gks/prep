package main

import (
    "fmt"
    "sync"
    "time"
	"errors"
)

var (
    ErrParkingFull   = errors.New("parking lot is full")
    ErrInvalidTicket = errors.New("ticket is invalid or already used")
)

// ParkingLot holds floors and slots
type ParkingLot struct {
    floors       int
    slotsPerFloor int
    slots        map[int]bool      // slotID -> occupied
    tickets      map[string]*Ticket
    mu           sync.Mutex
    nextTicket   int
}

// NewParkingLot ctor
func NewParkingLot(floors, slotsPerFloor int) *ParkingLot {
    total := floors * slotsPerFloor
    slots := make(map[int]bool, total)
    for i := 1; i <= total; i++ {
        slots[i] = false
    }
    return &ParkingLot{
        floors:       floors,
        slotsPerFloor: slotsPerFloor,
        slots:        slots,
        tickets:      make(map[string]*Ticket),
    }
}

// Park a vehicle, returns a ticket
func (p *ParkingLot) Park(v Vehicle) (*Ticket, error) {
    p.mu.Lock()
    defer p.mu.Unlock()

    for id, occupied := range p.slots {
        if !occupied {
            p.slots[id] = true
            ticketID := p.generateTicketID(v.GetNumber())
            ticket := &Ticket{ID: ticketID, VehicleNum: v.GetNumber(), SlotID: id}
            p.tickets[ticketID] = ticket
            Infof("Parked vehicle %s at slot %d", v.GetNumber(), id)
            return ticket, nil
        }
    }
    return nil, ErrParkingFull
}

// Unpark removes vehicle by ticket
func (p *ParkingLot) Unpark(ticketID string) error {
    p.mu.Lock()
    defer p.mu.Unlock()

    ticket, ok := p.tickets[ticketID]
    if !ok {
        return ErrInvalidTicket
    }
    delete(p.tickets, ticketID)
    p.slots[ticket.SlotID] = false
    Infof("Unparked vehicle %s from slot %d", ticket.VehicleNum, ticket.SlotID)
    return nil
}

func (p *ParkingLot) generateTicketID(vehicleNum string) string {
    p.nextTicket++
    return fmt.Sprintf("TKT-%s-%d", vehicleNum, time.Now().UnixNano()+int64(p.nextTicket))
}