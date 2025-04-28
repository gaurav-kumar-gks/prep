package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

// Booking represents a seat reservation
type Booking struct {
    ID       string
    ScreenID string
    Seat     SeatID
}

// BookingService handles booking logic
type BookingService struct {
    screens  map[string]*Screen
    bookings map[string]*Booking
    mu       sync.Mutex    // protects screens & bookings maps
    nextID   uint64        // atomic counter for booking IDs
}

// NewBookingService constructs the service
func NewBookingService() *BookingService {
    return &BookingService{
        screens:  make(map[string]*Screen),
        bookings: make(map[string]*Booking),
    }
}

// AddScreen registers a new screen in the system
func (bs *BookingService) AddScreen(id string, rows, cols int) {
    bs.mu.Lock()
    defer bs.mu.Unlock()
    bs.screens[id] = NewScreen(id, rows, cols)
}

// BookSeat reserves a seat, returns booking ID or error
func (bs *BookingService) BookSeat(screenID string, row, number int) (string, error) {
    bs.mu.Lock()
    screen, ok := bs.screens[screenID]
    bs.mu.Unlock()
    if !ok {
        return "", ErrScreenNotFound
    }
    if row < 1 || row > screen.Rows || number < 1 || number > screen.Cols {
        return "", ErrInvalidSeat
    }
    seatKey := SeatID{row, number}.String()
    screen.mu.Lock()
    defer screen.mu.Unlock()
    if screen.seats[seatKey] {
        return "", ErrSeatAlreadyBooked
    }
    // mark as booked
    screen.seats[seatKey] = true
    // generate booking ID
    id := fmt.Sprintf("BKG-%d", atomic.AddUint64(&bs.nextID, 1))
    booking := &Booking{ID: id, ScreenID: screenID, Seat: SeatID{row, number}}
    bs.mu.Lock()
    bs.bookings[id] = booking
    bs.mu.Unlock()
    return id, nil
}

// CancelBooking frees a seat by booking ID
func (bs *BookingService) CancelBooking(bookingID string) error {
    bs.mu.Lock()
    booking, ok := bs.bookings[bookingID]
    if !ok {
        bs.mu.Unlock()
        return ErrBookingNotFound
    }
    delete(bs.bookings, bookingID)
    bs.mu.Unlock()

    bs.mu.Lock()
    screen, exists := bs.screens[booking.ScreenID]
    bs.mu.Unlock()
    if !exists {
        return ErrScreenNotFound
    }
    key := booking.Seat.String()
    screen.mu.Lock()
    defer screen.mu.Unlock()
    screen.seats[key] = false
    return nil
}

// GetAvailableSeats returns all unbooked seats for a screen
func (bs *BookingService) GetAvailableSeats(screenID string) ([]SeatID, error) {
    bs.mu.Lock()
    screen, ok := bs.screens[screenID]
    bs.mu.Unlock()
    if !ok {
        return nil, ErrScreenNotFound
    }
    screen.mu.Lock()
    defer screen.mu.Unlock()
    avail := make([]SeatID, 0)
    for r := 1; r <= screen.Rows; r++ {
        for c := 1; c <= screen.Cols; c++ {
            key := SeatID{r, c}.String()
            if !screen.seats[key] {
                avail = append(avail, SeatID{r, c})
            }
        }
    }
    return avail, nil
}