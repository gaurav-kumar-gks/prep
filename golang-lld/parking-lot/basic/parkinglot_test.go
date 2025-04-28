package main

import (
    "testing"
)

func TestPark_Unpark(t *testing.T) {
    pl := NewParkingLot(1, 2)
    car1 := NewCar("KA-01-1111")
    car2 := NewCar("KA-01-2222")
    ticket1, err := pl.Park(car1)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
    _, park2err := pl.Park(car2)
    if park2err != nil {
        t.Fatalf("expected no error, got %v", park2err)
    }
    _, err = pl.Park(NewCar("KA-01-3333"))
    if err != ErrParkingFull {
        t.Fatalf("expected ErrParkingFull, got %v", err)
    }
    if err := pl.Unpark(ticket1.ID); err != nil {
        t.Fatalf("unexpected error on unpark: %v", err)
    }
    // Now one slot free
    _, err = pl.Park(NewCar("KA-01-3333"))
    if err != nil {
        t.Fatalf("expected success after free, got %v", err)
    }
}
