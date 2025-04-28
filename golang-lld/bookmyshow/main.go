package main

import (
    "fmt"
    "log"
)

func main() {
    svc := NewBookingService()
    svc.AddScreen("Screen1", 3, 5)

    // Book a seat
    bkgID, err := svc.BookSeat("Screen1", 2, 4)
    if err != nil {
        log.Fatalf("booking failed: %v", err)
    }
    fmt.Printf("Booked successfully, ID: %s\n", bkgID)

    // Attempt double-book on same seat
    if _, err := svc.BookSeat("Screen1", 2, 4); err != ErrSeatAlreadyBooked {
        log.Fatalf("expected ErrSeatAlreadyBooked, got %v", err)
    }

    // List available seats
    avail, _ := svc.GetAvailableSeats("Screen1")
    fmt.Println("Available seats:", avail)

    // Cancel booking
    if err := svc.CancelBooking(bkgID); err != nil {
        log.Fatalf("cancel failed: %v", err)
    }
    fmt.Println("Cancellation successful")

    // Now the seat is available again
    avail, _ = svc.GetAvailableSeats("Screen1")
    fmt.Println("Available seats after cancel:", avail)
}
