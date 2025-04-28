package main

import "errors"

var (
    ErrInvalidNotification   = errors.New("invalid notification")
    ErrMaxRetriesExceeded    = errors.New("max retries exceeded")
)