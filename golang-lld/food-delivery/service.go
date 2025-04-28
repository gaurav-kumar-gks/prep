package main

import (
    "fmt"
    "math"
    "sync"
    "sync/atomic"
)

// FoodDeliveryService handles core operations
type FoodDeliveryService struct {
    mu             sync.Mutex
    restaurants    map[string]*Restaurant
    orders         map[string]*Order
    agents         map[string]*DeliveryAgent
    nextOrderID    uint64
}

// NewService initializes the service
func NewService() *FoodDeliveryService {
    return &FoodDeliveryService{
        restaurants: make(map[string]*Restaurant),
        orders:      make(map[string]*Order),
        agents:      make(map[string]*DeliveryAgent),
    }
}

// RegisterRestaurant adds a new restaurant
func (s *FoodDeliveryService) RegisterRestaurant(r Restaurant) {
    s.mu.Lock()
    defer s.mu.Unlock()
    if r.Menu == nil {
        r.Menu = make(map[string]Item)
    }
    s.restaurants[r.ID] = &r
}

// AddMenuItem adds/updates an item in restaurant menu
func (s *FoodDeliveryService) AddMenuItem(restaurantID string, item Item) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    r, ok := s.restaurants[restaurantID]
    if !ok {
        return ErrRestaurantNotFound
    }
    r.Menu[item.ID] = item
    return nil
}

// RegisterAgent adds a delivery agent
func (s *FoodDeliveryService) RegisterAgent(agent DeliveryAgent) {
    s.mu.Lock()
    defer s.mu.Unlock()
    agent.Available = true
    s.agents[agent.ID] = &agent
}

// PlaceOrder creates an order and returns its ID
func (s *FoodDeliveryService) PlaceOrder(restaurantID string, itemIDs []string) (string, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    r, ok := s.restaurants[restaurantID]
    if !ok {
        return "", ErrRestaurantNotFound
    }
    for _, id := range itemIDs {
        if _, exists := r.Menu[id]; !exists {
            return "", ErrItemNotInMenu
        }
    }
    orderID := fmt.Sprintf("ORD-%d", atomic.AddUint64(&s.nextOrderID, 1))
    order := &Order{ID: orderID, RestaurantID: restaurantID, ItemIDs: itemIDs, Status: StatusPlaced}
    s.orders[orderID] = order
    return orderID, nil
}

// ConfirmOrder updates status to CONFIRMED
func (s *FoodDeliveryService) ConfirmOrder(orderID string) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    o, ok := s.orders[orderID]
    if !ok {
        return ErrOrderNotFound
    }
    o.Status = StatusConfirmed
    return nil
}

// PrepareOrder updates status to PREPARED
func (s *FoodDeliveryService) PrepareOrder(orderID string) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    o, ok := s.orders[orderID]
    if !ok {
        return ErrOrderNotFound
    }
    o.Status = StatusPrepared
    return nil
}

// AssignAgent assigns nearest available agent to prepared order
func (s *FoodDeliveryService) AssignAgent(orderID string) (string, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    o, ok := s.orders[orderID]
    if !ok {
        return "", ErrOrderNotFound
    }
    if o.Status != StatusPrepared {
        return "", fmt.Errorf("order not prepared")
    }
    r := s.restaurants[o.RestaurantID]
    bestID := ""
    bestDist := math.MaxFloat64
    for id, ag := range s.agents {
        if !ag.Available {
            continue
        }
        dist := haversine(ag.Location, r.Location)
        if dist < bestDist {
            bestDist, bestID = dist, id
        }
    }
    if bestID == "" {
        return "", ErrNoDeliveryAgent
    }
    agent := s.agents[bestID]
    agent.Available = false
    o.DeliveryAgentID = bestID
    o.Status = StatusOutForDelivery
    return bestID, nil
}

// CompleteDelivery marks order delivered and frees agent
func (s *FoodDeliveryService) CompleteDelivery(orderID string) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    o, ok := s.orders[orderID]
    if !ok {
        return ErrOrderNotFound
    }
    ag, ok := s.agents[o.DeliveryAgentID]
    if !ok {
        return ErrInvalidAgent
    }
    o.Status = StatusDelivered
    ag.Available = true
    return nil
}

// GetOrder returns order details
func (s *FoodDeliveryService) GetOrder(orderID string) (*Order, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    o, ok := s.orders[orderID]
    if !ok {
        return nil, ErrOrderNotFound
    }
    return o, nil
}

func haversine(a, b Location) float64 {
    const R = 6371e3 // Earth radius in meters
    toRad := func(deg float64) float64 { return deg * math.Pi / 180 }
    φ1, φ2 := toRad(a.Lat), toRad(b.Lat)
    Δφ := toRad(b.Lat - a.Lat)
    Δλ := toRad(b.Lon - a.Lon)

    sinΔφ := math.Sin(Δφ/2)
    sinΔλ := math.Sin(Δλ/2)

    h := sinΔφ*sinΔφ + math.Cos(φ1)*math.Cos(φ2)*sinΔλ*sinΔλ
    c := 2 * math.Atan2(math.Sqrt(h), math.Sqrt(1-h))
    return R * c
}