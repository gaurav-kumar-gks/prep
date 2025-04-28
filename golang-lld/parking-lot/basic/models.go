package main

// VehicleType represents the type of vehicle
type VehicleType string

const (
    CarType   VehicleType = "Car"
    BikeType  VehicleType = "Bike"
    TruckType VehicleType = "Truck"
)

// Vehicle defines behavior
type Vehicle interface {
    GetType() VehicleType
    GetNumber() string
}

// Car struct
type Car struct { number string }
func NewCar(num string) *Car { return &Car{number: num} }
func (c *Car) GetType() VehicleType { return CarType }
func (c *Car) GetNumber() string    { return c.number }

// Ticket issued on parking
type Ticket struct {
    ID         string
    VehicleNum string
    SlotID     int
}
