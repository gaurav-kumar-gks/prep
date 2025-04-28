package main

type VehicleType string

const (
    CarType   VehicleType = "Car"
    BikeType  VehicleType = "Bike"
    TruckType VehicleType = "Truck"
)

type Vehicle interface {
    GetType() VehicleType
    GetNumber() string
}

type Car struct{ Number string }
func NewCar(num string) *Car { return &Car{Number: num} }
func (c *Car) GetType() VehicleType { return CarType }
func (c *Car) GetNumber() string    { return c.Number }

type Ticket struct {
    ID         string      `json:"id"`
    VehicleNum string      `json:"vehicle_num"`
    SlotID     int         `json:"slot_id"`
}

type APIResponse struct {
    Data    interface{}   `json:"data,omitempty"`
    Message string        `json:"message,omitempty"`
    Errors  []string      `json:"errors,omitempty"`
}