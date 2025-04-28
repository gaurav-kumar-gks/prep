package main

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *sql.DB {
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        t.Fatalf("failed to open test DB: %v", err)
    }

    // Create tables with the correct schema
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS slots (
            slot_id INTEGER PRIMARY KEY,
            occupied BOOLEAN DEFAULT FALSE
        );
        CREATE TABLE IF NOT EXISTS tickets (
            id TEXT PRIMARY KEY,
            vehicle_num TEXT,
            slot_id INTEGER,
            FOREIGN KEY(slot_id) REFERENCES slots(slot_id)
        );
    `)
    if err != nil {
        t.Fatalf("failed to create tables: %v", err)
    }

    // Insert some test slots
    _, err = db.Exec("INSERT INTO slots (slot_id, occupied) VALUES (1, 0), (2, 0)")
    if err != nil {
        t.Fatalf("failed to insert test slots: %v", err)
    }

    return db
}

func TestRepo_FindAndOccupySlot(t *testing.T) {
    db := setupTestDB(t)
    defer db.Close()

    repo := NewRepo(db)
    tx, _ := db.Begin()
	defer tx.Rollback()
    slotID, err := repo.FindAndOccupySlot(tx)
    if err != nil || slotID!=1 {
        t.Fatalf("unexpected: %v, %d", err, slotID)
    }
}
