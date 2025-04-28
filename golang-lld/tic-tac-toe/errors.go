package main

import "errors"

var (
    ErrGameNotFound  = errors.New("game not found")
    ErrOutOfBounds   = errors.New("move out of bounds")
    ErrCellOccupied  = errors.New("cell already occupied")
    ErrGameEnded     = errors.New("game already ended")
)