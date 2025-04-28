package main

import "errors"

var (
    ErrScreenNotFound    = errors.New("screen not found")
    ErrSeatAlreadyBooked = errors.New("seat already booked")
    ErrInvalidSeat       = errors.New("invalid seat coordinates")
    ErrBookingNotFound   = errors.New("booking not found")
)