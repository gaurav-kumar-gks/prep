package main

import (
    "math"
    "sync"
)

// CabService manages drivers and ride requests
type CabService struct {
    mu       sync.Mutex
    drivers  map[string]*Driver
    requests map[string]*RideRequest
}

// NewCabService creates a new service instance
func NewCabService() *CabService {
    return &CabService{
        drivers:  make(map[string]*Driver),
        requests: make(map[string]*RideRequest),
    }
}

// RegisterDriver adds a driver to the system
func (cs *CabService) RegisterDriver(id string, loc Location) {
    cs.mu.Lock()
    defer cs.mu.Unlock()
    cs.drivers[id] = &Driver{ID: id, Location: loc, Status: Available}
}

// UpdateDriverLocation updates a driver's location
func (cs *CabService) UpdateDriverLocation(id string, loc Location) error {
    cs.mu.Lock()
    defer cs.mu.Unlock()
    d, ok := cs.drivers[id]
    if !ok {
        return ErrInvalidDriver
    }
    d.Location = loc
    return nil
}

// RequestRide handles a ride request and returns assigned driver ID or error
func (cs *CabService) RequestRide(reqID string, pickup, dropoff Location) (string, error) {
    cs.mu.Lock()
    defer cs.mu.Unlock()
    // find nearest available driver
    bestID := ""
    bestDist := math.MaxFloat64
    for id, d := range cs.drivers {
        if d.Status != Available {
            continue
        }
        dist := haversine(d.Location, pickup)
        if dist < bestDist {
            bestDist, bestID = dist, id
        }
    }
    if bestID == "" {
        return "", ErrNoDriverAvailable
    }
    // assign
    cs.drivers[bestID].Status = Busy
    request := &RideRequest{ID: reqID, Pickup: pickup, Dropoff: dropoff, Status: Assigned, DriverID: bestID}
    cs.requests[reqID] = request
    return bestID, nil
}

// CompleteRide marks a ride as completed and frees driver
func (cs *CabService) CompleteRide(reqID string) error {
    cs.mu.Lock()
    defer cs.mu.Unlock()
    r, ok := cs.requests[reqID]
    if !ok {
        return ErrInvalidRequest
    }
    d, ok := cs.drivers[r.DriverID]
    if !ok {
        return ErrInvalidDriver
    }
    r.Status = Completed
    d.Status = Available
    return nil
}

// GetRequest returns current ride request state
func (cs *CabService) GetRequest(reqID string) (*RideRequest, error) {
    cs.mu.Lock()
    defer cs.mu.Unlock()
    r, ok := cs.requests[reqID]
    if !ok {
        return nil, ErrInvalidRequest
    }
    return r, nil
}

// haversine computes great-circle distance between two locations
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