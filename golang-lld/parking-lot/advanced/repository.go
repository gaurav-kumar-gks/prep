package main

import (
    "database/sql"

)

// ParkingLotRepo defines DB operations
type ParkingLotRepo interface {
    FindAndOccupySlot(tx *sql.Tx) (int, error)
    CreateTicket(tx *sql.Tx, ticket *Ticket) error
    GetTicketSlot(tx *sql.Tx, ticketID string) (int, error)
    DeleteTicket(tx *sql.Tx, ticketID string) error
    FreeSlot(tx *sql.Tx, slotID int) error
}

// repo implements ParkingLotRepo using SQLite
type repo struct {
    db *sql.DB
}

func NewRepo(db *sql.DB) ParkingLotRepo {
    return &repo{db: db}
}

func (r *repo) FindAndOccupySlot(tx *sql.Tx) (int, error) {
    row := tx.QueryRow(`SELECT slot_id FROM slots WHERE occupied=0 LIMIT 1`)
    var slotID int
    if err := row.Scan(&slotID); err != nil {
        return 0, err
    }
    if _, err := tx.Exec(`UPDATE slots SET occupied=1 WHERE slot_id=?`, slotID); err != nil {
        return 0, err
    }
    return slotID, nil
}

func (r *repo) CreateTicket(tx *sql.Tx, ticket *Ticket) error {
    _, err := tx.Exec(`INSERT INTO tickets(id, vehicle_num, slot_id) VALUES(?,?,?)`,
        ticket.ID, ticket.VehicleNum, ticket.SlotID)
    return err
}

func (r *repo) GetTicketSlot(tx *sql.Tx, ticketID string) (int, error) {
    row := tx.QueryRow(`SELECT slot_id FROM tickets WHERE id=?`, ticketID)
    var slotID int
    err := row.Scan(&slotID)
    return slotID, err
}

func (r *repo) DeleteTicket(tx *sql.Tx, ticketID string) error {
    _, err := tx.Exec(`DELETE FROM tickets WHERE id=?`, ticketID)
    return err
}

func (r *repo) FreeSlot(tx *sql.Tx, slotID int) error {
    _, err := tx.Exec(`UPDATE slots SET occupied=0 WHERE slot_id=?`, slotID)
    return err
}