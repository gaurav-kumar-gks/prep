package main

// Location struct for restaurants and agents
type Location struct { Lat, Lon float64 }

// Item represents a menu item with price
type Item struct {
    ID    string
    Name  string
    Price float64
}

// Restaurant holds menu and info
type Restaurant struct {
    ID       string
    Name     string
    Location Location
    Menu     map[string]Item // itemID -> Item
}

// OrderStatus enumeration
type OrderStatus string

const (
    StatusPlaced        OrderStatus = "PLACED"
    StatusConfirmed     OrderStatus = "CONFIRMED"
    StatusPrepared      OrderStatus = "PREPARED"
    StatusOutForDelivery OrderStatus = "OUT_FOR_DELIVERY"
    StatusDelivered     OrderStatus = "DELIVERED"
    StatusCancelled     OrderStatus = "CANCELLED"
)

// Order captures customer order
type Order struct {
    ID           string
    RestaurantID string
    ItemIDs      []string
    Status       OrderStatus
    DeliveryAgentID string
}

// DeliveryAgent represents a courier
type DeliveryAgent struct {
    ID       string
    Location Location
    Available bool
}