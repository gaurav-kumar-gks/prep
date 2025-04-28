package main

import (
	"database/sql"
	"errors"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// mockRepo for testing service
type mockRepo struct {
    slotID int
    err    error
}
func (m *mockRepo) FindAndOccupySlot(tx *sql.Tx) (int, error) { return m.slotID, m.err }
func (m *mockRepo) CreateTicket(tx *sql.Tx, ticket *Ticket) error { return m.err }
func (m *mockRepo) GetTicketSlot(tx *sql.Tx, ticketID string) (int, error) { if m.err!=nil {return 0,m.err}; return m.slotID,nil }
func (m *mockRepo) DeleteTicket(tx *sql.Tx, ticketID string) error { return m.err }
func (m *mockRepo) FreeSlot(tx *sql.Tx, slotID int) error { return m.err }

// testSetup creates a test environment with an in-memory database
func testSetup(t *testing.T) (*sql.DB, *ParkingLotService) {
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        t.Fatalf("failed to open test DB: %v", err)
    }
    
    // Create tables
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS slots (
            slot_id INTEGER PRIMARY KEY,
            occupied BOOLEAN DEFAULT FALSE
        );
        CREATE TABLE IF NOT EXISTS tickets (
            id TEXT PRIMARY KEY,
            slot_id INTEGER,
            FOREIGN KEY(slot_id) REFERENCES slots(slot_id)
        );
    `)
    if err != nil {
        t.Fatalf("failed to create tables: %v", err)
    }
    
    return db, NewService(NewRepo(db), 1, 1)
}

func TestService_Park_Fail(t *testing.T) {
    db, svc := testSetup(t)
    defer db.Close()
    
    // Override the repository with a mock that returns an error
    svc.repo = &mockRepo{err: errors.New("no slots")}
    
    _, err := svc.Park(db, NewCar("X"))
    if err != ErrParkingFull {
        t.Fatalf("expected ErrParkingFull, got %v", err)
    }
}

func TestService_Unpark_Fail(t *testing.T) {
    db, svc := testSetup(t)
    defer db.Close()
    
    // Override the repository with a mock that returns an error
    svc.repo = &mockRepo{err: errors.New("not found")}
    
    err := svc.Unpark(db, "T1")
    if err != ErrInvalidTicket {
        t.Fatalf("expected ErrInvalidTicket, got %v", err)
    }
}