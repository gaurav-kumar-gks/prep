package main

import "errors"

var (
    ErrRestaurantNotFound   = errors.New("restaurant not found")
    ErrItemNotInMenu        = errors.New("item not in menu")
    ErrOrderNotFound        = errors.New("order not found")
    ErrNoDeliveryAgent      = errors.New("no available delivery agent")
    ErrInvalidAgent         = errors.New("invalid delivery agent ID")
)
