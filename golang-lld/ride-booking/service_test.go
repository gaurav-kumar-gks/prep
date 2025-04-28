package main

import (
    "testing"
)

func TestRideFlow(t *testing.T) {
    svc := NewCabService()
    svc.RegisterDriver("D1", Location{0,0})
    svc.RegisterDriver("D2", Location{10,10})
    // request near D1
    did, err := svc.RequestRide("R1", Location{1,1}, Location{2,2})
    if err != nil || did != "D1" {
        t.Errorf("expected D1, got %s, err %v", did, err)
    }
    // second request should pick D2
    did2, err := svc.RequestRide("R2", Location{1,1}, Location{2,2})
    if err != nil || did2 != "D2" {
        t.Errorf("expected D2, got %s, err %v", did2, err)
    }
    // no more drivers
    if _, err := svc.RequestRide("R3", Location{1,1}, Location{2,2}); err != ErrNoDriverAvailable {
        t.Errorf("expected ErrNoDriverAvailable, got %v", err)
    }
    // complete R1
    if err := svc.CompleteRide("R1"); err != nil {
        t.Errorf("complete failed: %v", err)
    }
    // now D1 available for new ride
    did3, err := svc.RequestRide("R3", Location{1,1}, Location{2,2})
    if err != nil || did3 != "D1" {
        t.Errorf("expected D1, got %s, err %v", did3, err)
    }
}
