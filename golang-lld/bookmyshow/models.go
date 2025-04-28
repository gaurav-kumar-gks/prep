package main

import (
    "fmt"
    "sync"
)

// SeatID uniquely identifies a seat within a screen
type SeatID struct {
    Row    int
    Number int
}

func (s SeatID) String() string {
    return fmt.Sprintf("R%dS%d", s.Row, s.Number)
}

// Screen represents a single auditorium or screen with seats
type Screen struct {
    ID    string
    Rows  int
    Cols  int
    seats map[string]bool // seatID string -> booked status
    mu    sync.Mutex      // protects seats map
}

func NewScreen(id string, rows, cols int) *Screen {
    total := rows * cols
    m := make(map[string]bool, total)
    for r := 1; r <= rows; r++ {
        for c := 1; c <= cols; c++ {
            m[SeatID{r, c}.String()] = false
        }
    }
    return &Screen{ID: id, Rows: rows, Cols: cols, seats: m}
}