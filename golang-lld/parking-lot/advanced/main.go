package main

import (
	"database/sql"
	"log"
	"net/http"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
    // README.md has local DB setup instructions

    // db, err := sql.Open("sqlite3", "./parking.db")
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        log.Fatalf("failed to open DB: %v", err)
    }
    defer db.Close()
    // pre-create tables via init_db.sql or manual step

	// create tables
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS slots (
            slot_id INTEGER PRIMARY KEY,
            occupied BOOLEAN DEFAULT FALSE
        );
        CREATE TABLE IF NOT EXISTS tickets (
            id TEXT PRIMARY KEY,
            vehicle_num TEXT,
            slot_id INTEGER,
            FOREIGN KEY(slot_id) REFERENCES slots(slot_id)
        );
	`)
	if err != nil {
		log.Fatalf("failed to create tables: %v", err)
	}
    
    repo := NewRepo(db)
    svc := NewService(repo, 2, 5)

    handlers := map[string]http.HandlerFunc{
        "/park":   methodHandler(http.MethodPost, ParkHandler(svc, db)),
        "/unpark": methodHandler(http.MethodPost, UnparkHandler(svc, db)),
    }
    router := NewRouter(handlers)

    log.Println("Server listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func NewRouter(handlerFuncs map[string]http.HandlerFunc) http.Handler {
    mux := http.NewServeMux()
    for path, handler := range handlerFuncs {
        mux.HandleFunc(path, handler)
    }
    return mux
}

func methodHandler(method string, handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != method {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }
        handler(w, r)
    }
}

func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Header.Get("X-API-KEY") != "secret-key" {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte(`{"errors":["unauthorized"]}`))
            return
        }
        next.ServeHTTP(w, r)
    })
}

func rateLimiterMiddleware(rps int) func(http.Handler) http.Handler {
    var mu sync.Mutex
    tokens := rps

    go func() {
        ticker := time.NewTicker(time.Second)
        for range ticker.C {
            mu.Lock()
            tokens = rps
            mu.Unlock()
        }
    }()

    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            mu.Lock()
            defer mu.Unlock()
            if tokens <= 0 {
                w.WriteHeader(http.StatusTooManyRequests)
                w.Write([]byte(`{"errors":["rate limit exceeded"]}`))
                return
            }
            tokens--
            next.ServeHTTP(w, r)
        })
    }
}
