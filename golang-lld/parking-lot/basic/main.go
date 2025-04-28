package main

import (
    "fmt"
	"log"
)

func Infof(fmtStr string, args ...interface{}) {
    log.Printf("[INFO] "+fmtStr, args...)
}

func Errorf(fmtStr string, args ...interface{}) {
    log.Printf("[ERROR] "+fmtStr, args...)
}

func main() {
    parkingLot := NewParkingLot(2, 5)
    car := NewCar("KA-01-1234")

    ticket, err := parkingLot.Park(car)
    if err != nil {
        Errorf("Failed to park: %v", err)
        return
    }
    fmt.Printf("Parked %s, ticket ID: %s\n", car.GetNumber(), ticket.ID)

    err = parkingLot.Unpark(ticket.ID)
    if err != nil {
        Errorf("Failed to unpark: %v", err)
        return
    }
    fmt.Println("Unparked successfully")
}
