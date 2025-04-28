package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// ParkHandler handles /park endpoint
func ParkHandler(svc *ParkingLotService, db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        
        var req struct{ Number string `json:"number"` }
        var resp APIResponse
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            w.WriteHeader(http.StatusBadRequest)
            resp.Errors = []string{err.Error()}
            json.NewEncoder(w).Encode(resp)
            return
        }
        ticket, err := svc.Park(db, NewCar(req.Number))
        if err != nil {
            w.WriteHeader(http.StatusConflict)
            resp.Errors = []string{err.Error()}
            json.NewEncoder(w).Encode(resp)
            return
        }
        w.WriteHeader(http.StatusCreated)
        resp.Data = ticket
        resp.Message = "parked successfully"
        json.NewEncoder(w).Encode(resp)
    }
}

// UnparkHandler handles /unpark endpoint
func UnparkHandler(svc *ParkingLotService, db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        
        ticketID := r.URL.Query().Get("ticket_id")
        var resp APIResponse
        if ticketID == "" {
            w.WriteHeader(http.StatusBadRequest)
            resp.Errors = []string{"ticket_id is required"}
            json.NewEncoder(w).Encode(resp)
            return
        }
        if err := svc.Unpark(db, ticketID); err != nil {
            w.WriteHeader(http.StatusNotFound)
            resp.Errors = []string{err.Error()}
            json.NewEncoder(w).Encode(resp)
            return
        }
        w.WriteHeader(http.StatusOK)
        resp.Message = "unparked successfully"
        json.NewEncoder(w).Encode(resp)
    }
}