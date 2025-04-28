package main

import (
    "testing"
)

func TestCompleteFlow(t *testing.T) {
    svc := NewService()
    // setup restaurant and items
    svc.RegisterRestaurant(Restaurant{ID: "R1", Name: "PizzaPlace", Location: Location{0,0}})
    svc.AddMenuItem("R1", Item{ID: "I1", Name: "Margherita", Price: 5.0})
    svc.AddMenuItem("R1", Item{ID: "I2", Name: "Pepperoni", Price: 7.0})
    // register agents
    svc.RegisterAgent(DeliveryAgent{ID: "A1", Location: Location{0,1}})
    svc.RegisterAgent(DeliveryAgent{ID: "A2", Location: Location{10,10}})

    // place order
    oid, err := svc.PlaceOrder("R1", []string{"I1","I2"})
    if err != nil {
        t.Fatalf("place order failed: %v", err)
    }
    if err := svc.ConfirmOrder(oid); err != nil {
        t.Fatalf("confirm failed: %v", err)
    }
    if err := svc.PrepareOrder(oid); err != nil {
        t.Fatalf("prepare failed: %v", err)
    }
    agID, err := svc.AssignAgent(oid)
    if err != nil {
        t.Fatalf("assign agent failed: %v", err)
    }
    if agID != "A1" {
        t.Errorf("expected A1, got %s", agID)
    }
    if err := svc.CompleteDelivery(oid); err != nil {
        t.Fatalf("complete delivery failed: %v", err)
    }
    // after completion, agent should be available
    order, _ := svc.GetOrder(oid)
    if order.Status != StatusDelivered {
        t.Errorf("expected StatusDelivered, got %s", order.Status)
    }
    ag, _ := svc.agents[agID]
    if !ag.Available {
        t.Errorf("expected agent available after delivery")
    }
}
