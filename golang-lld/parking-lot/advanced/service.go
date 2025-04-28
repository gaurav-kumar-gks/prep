package main

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"
)

var (
    ErrParkingFull   = errors.New("parking lot is full")
    ErrInvalidTicket = errors.New("ticket is invalid or already used")
)

// ParkingLotService defines business logic, decoupled from DB
type ParkingLotService struct {
    repo          ParkingLotRepo
    floors        int
    slotsPerFloor int
    mu            sync.Mutex  // ensure thread-safety if DB isolation not sufficient
}

// NewService creates a service with injected repository
func NewService(repo ParkingLotRepo, floors, slotsPerFloor int) *ParkingLotService {
    ps := &ParkingLotService{repo: repo, floors: floors, slotsPerFloor: slotsPerFloor}
    ps.populateSlots()
	return ps
}		

func (svc *ParkingLotService) populateSlots() {
	repo := svc.repo.(*repo)
	for i := 0; i < svc.floors * svc.slotsPerFloor; i++ {
		repo.db.Exec(`
			INSERT INTO slots (slot_id, occupied) VALUES (?, ?);
		`, i, false)
	}
}

func (svc *ParkingLotService) Park(db *sql.DB, v Vehicle) (*Ticket, error) {
    svc.mu.Lock()
    defer svc.mu.Unlock()
    
    tx, err := db.Begin()
    if err != nil {
        return nil, err
    }
    
    committed := false
    defer func() {
        if !committed {
            tx.Rollback()
        }
    }()

    slotID, err := svc.repo.FindAndOccupySlot(tx)
    if err != nil {
        return nil, ErrParkingFull
    }
    
    ticketID := fmt.Sprintf("TKT-%s-%d", v.GetNumber(), slotID)
    ticket := &Ticket{ID: ticketID, VehicleNum: v.GetNumber(), SlotID: slotID}
    
    if err := svc.repo.CreateTicket(tx, ticket); err != nil {
        return nil, err
    }
    
    if err := tx.Commit(); err != nil {
        return nil, err
    }
    
    committed = true
    return ticket, nil
}

func (svc *ParkingLotService) Unpark(db *sql.DB, ticketID string) error {
    svc.mu.Lock()
    defer svc.mu.Unlock()
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    committed := false
    defer func() {
        if !committed {
            tx.Rollback()
        }
    }()

    slotID, err := svc.repo.GetTicketSlot(tx, ticketID)
    if err != nil {
        return ErrInvalidTicket
    }
    if err := svc.repo.DeleteTicket(tx, ticketID); err != nil {
        return err
    }
    if err := svc.repo.FreeSlot(tx, slotID); err != nil {
        return err
    }
    if err := tx.Commit(); err != nil {
		return err
	}
	committed = true
	return nil
}