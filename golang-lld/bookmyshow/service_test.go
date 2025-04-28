package main

import (
    "sync"
    "testing"
)

func TestAddAndGetAvailableSeats(t *testing.T) {
    svc := NewBookingService()
    svc.AddScreen("S1", 2, 2)

    seats, err := svc.GetAvailableSeats("S1")
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if len(seats) != 4 {
        t.Errorf("expected 4 seats, got %d", len(seats))
    }
}

func TestInvalidScreen(t *testing.T) {
    svc := NewBookingService()
    if _, err := svc.GetAvailableSeats("NoScreen"); err != ErrScreenNotFound {
        t.Errorf("expected ErrScreenNotFound, got %v", err)
    }
    if _, err := svc.BookSeat("NoScreen", 1, 1); err != ErrScreenNotFound {
        t.Errorf("expected ErrScreenNotFound, got %v", err)
    }
    if err := svc.CancelBooking("BKG-999"); err != ErrBookingNotFound {
        t.Errorf("expected ErrBookingNotFound, got %v", err)
    }
}

func TestBookCancelFlow(t *testing.T) {
    svc := NewBookingService()
    svc.AddScreen("S2", 1, 1)

    id, err := svc.BookSeat("S2", 1, 1)
    if err != nil {
        t.Fatalf("booking failed: %v", err)
    }
    if _, err := svc.BookSeat("S2", 1, 1); err != ErrSeatAlreadyBooked {
        t.Errorf("expected ErrSeatAlreadyBooked, got %v", err)
    }
    if err := svc.CancelBooking(id); err != nil {
        t.Errorf("cancel failed: %v", err)
    }
    // after cancel, should succeed
    if _, err := svc.BookSeat("S2", 1, 1); err != nil {
        t.Errorf("expected success after cancel, got %v", err)
    }
}

func TestConcurrentBooking(t *testing.T) {
    svc := NewBookingService()
    svc.AddScreen("S3", 1, 10)
    var wg sync.WaitGroup
    successCount := 0
    mu := sync.Mutex{}

    for i := 1; i <= 10; i++ {
        wg.Add(1)
        go func(seatNum int) {
            defer wg.Done()
            if _, err := svc.BookSeat("S3", 1, seatNum); err == nil {
                mu.Lock()
                successCount++
                mu.Unlock()
            }
        }(i)
    }
    wg.Wait()

    if successCount != 10 {
        t.Errorf("expected 10 successful bookings, got %d", successCount)
    }
    seats, _ := svc.GetAvailableSeats("S3")
    if len(seats) != 0 {
        t.Errorf("expected 0 available seats, got %d", len(seats))
    }
}
