/*
LLD Design for Cab Booking System (Ride-Hailing) typically tests:
1) Data modelling of drivers, requests, and dynamic state transitions
2) Spatial matching algorithms (nearest-driver search) and basic load balancing
3) Thread-safe handling of driver location updates and concurrent ride requests

Typical Requirements:
- Register drivers with current location and status (available/busy)
- Update driver location in real-time
- Passenger requests ride with pickup & dropoff coordinates
- Match request to nearest available driver and mark driver busy
- Complete/cancel ride to free driver
- Query ride status and driver assignment

Typical Extensions & Solutions:
- **Surge Pricing**: adjust fare based on demand-supply ratio in zone; maintain zone-wise driver/request counts
- **Pool Rides**: allow matching multiple passengers along similar routes; solve via route overlap heuristics or graph-based clustering
- **Route Optimization**: use mapping APIs to compute ETA and select driver minimizing pickup time; incorporate traffic data
- **Geo-Partitioning**: partition service area into zones/quadtree for efficient nearest-neighbor search (e.g., KD-Tree or geohash-based indexing)
- **ETA Updates**: streaming ETA to passenger & driver using WebSockets, handling cancellations
- **Driver Stats & Rating**: persistence layer to record trip histories and ratings; analytics
- **High Availability**: shard service per region, use distributed cache (Redis) for quick lookups
*/

package main

import "fmt"

func main() {
    svc := NewCabService()
    // register drivers
    svc.RegisterDriver("D1", Location{Lat: 12.97, Lon: 77.59})
    svc.RegisterDriver("D2", Location{Lat: 12.98, Lon: 77.60})

    // passenger requests ride
    driverID, err := svc.RequestRide("R1", Location{12.96,77.58}, Location{12.99,77.61})
    if err != nil {
        fmt.Println("Ride request failed:", err)
        return
    }
    fmt.Println("Assigned Driver:", driverID)

    // complete ride
    if err := svc.CompleteRide("R1"); err != nil {
        fmt.Println("Complete ride failed:", err)
        return
    }
    fmt.Println("Ride completed for R1")
}
