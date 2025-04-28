/*
LLD Design for Food Delivery System (Zomato/Swiggy) typically tests:
1) Modeling of restaurants, menus, orders, and delivery agents
2) Order lifecycle management and state transitions
3) Matching and dispatching delivery agents (couriers) to orders
4) Concurrency-safe handling of menu updates and order placements

Typical Requirements:
- Register restaurants and their menus (items with price)
- Customers place orders: select restaurant and items
- Order status flows: Placed -> Confirmed -> Prepared -> OutForDelivery -> Delivered -> (optional) Cancelled
- Assign an available delivery agent when order is prepared
- Track delivery status and ETA (basic linear movement simulation)
- Query order details and status

Typical Extensions & Solutions:
- **Dynamic ETA**: incorporate real-time traffic data and courier GPS updates; use map APIs for route calculation
- **Surge/Dynamic Pricing**: increase delivery fee during peak times or high demand
- **Multi-restaurant Orders**: batch orders from nearby restaurants into a single delivery run
- **Load Balancing Agents**: shard couriers by zones or use weighted round-robin to distribute load
- **Promotions & Discounts**: apply coupon codes, loyalty points, and calculate final fare
- **Payment Integration**: handle multiple payment modes, wallets, refunds
- **High Availability**: split service by region, use distributed cache (Redis) for quick lookups
- **Persistence**: swap in repository layer for DB storage and recovery on restart
*/

package main

import "fmt"

func main() {
    svc := NewService()
    // Register restaurant and menu
    svc.RegisterRestaurant(Restaurant{ID: "R1", Name: "PizzaPlace", Location: Location{Lat: 12.97, Lon: 77.59}})
    svc.AddMenuItem("R1", Item{ID: "I1", Name: "Margherita", Price: 8.5})
    svc.AddMenuItem("R1", Item{ID: "I2", Name: "Farmhouse", Price: 10.0})

    // Register delivery agents
    svc.RegisterAgent(DeliveryAgent{ID: "A1", Location: Location{Lat: 12.98, Lon: 77.60}})
    svc.RegisterAgent(DeliveryAgent{ID: "A2", Location: Location{Lat: 12.96, Lon: 77.58}})

    // Place new order
    orderID, err := svc.PlaceOrder("R1", []string{"I1", "I2"})
    if err != nil {
        fmt.Println("Error placing order:", err)
        return
    }
    fmt.Println("Order placed:", orderID)

    // Progress order lifecycle
    _ = svc.ConfirmOrder(orderID)
    fmt.Println("Order confirmed")
    _ = svc.PrepareOrder(orderID)
    fmt.Println("Order prepared")

    // Assign a delivery agent
    agentID, err := svc.AssignAgent(orderID)
    if err != nil {
        fmt.Println("Error assigning agent:", err)
        return
    }
    fmt.Println("Agent assigned:", agentID)

    // Complete delivery
    _ = svc.CompleteDelivery(orderID)
    fmt.Println("Delivery completed for order:", orderID)
}