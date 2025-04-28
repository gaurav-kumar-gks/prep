package main

import "errors"

var (
    ErrNoDriverAvailable = errors.New("no available driver")
    ErrInvalidDriver     = errors.New("invalid driver ID")
    ErrInvalidRequest    = errors.New("invalid ride request ID")
)
