package main

// Location represents a geographic point
type Location struct {
    Lat float64
    Lon float64
}

// DriverStatus indicates availability
type DriverStatus string

const (
    Available DriverStatus = "AVAILABLE"
    Busy      DriverStatus = "BUSY"
)

// Driver captures driver state
type Driver struct {
    ID       string
    Location Location
    Status   DriverStatus
}

// RideStatus indicates current state of ride
type RideStatus string

const (
    Requested RideStatus = "REQUESTED"
    Assigned  RideStatus = "ASSIGNED"
    Completed RideStatus = "COMPLETED"
    Cancelled RideStatus = "CANCELLED"
)

// RideRequest captures a passenger request
type RideRequest struct {
    ID         string
    Pickup     Location
    Dropoff    Location
    Status     RideStatus
    DriverID   string
}