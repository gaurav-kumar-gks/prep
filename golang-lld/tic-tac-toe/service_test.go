package main

import (
    "sync"
    "testing"
)

func TestWinScenario(t *testing.T) {
    svc := NewGameService()
    id := svc.CreateGame(3)
    svc.MakeMove(id,1,1)
    svc.MakeMove(id,1,2)
    svc.MakeMove(id,2,2)
    svc.MakeMove(id,1,3)
    svc.MakeMove(id,3,3)
    game, _ := svc.GetGame(id)
    if game.State != StateWon || game.Winner != X {
        t.Errorf("expected X to win, got state=%s, winner=%s", game.State, game.Winner)
    }
}

func TestDrawScenario(t *testing.T) {
    svc := NewGameService()
    id := svc.CreateGame(3)
    seq := [][2]int{{1,1},{1,2},{1,3},{2,1},{2,3},{2,2},{3,1},{3,3},{3,2}}
    for _, mv := range seq {
        if err := svc.MakeMove(id, mv[0], mv[1]); err != nil {
            t.Fatalf("move %v failed: %v", mv, err)
        }
    }
    game, _ := svc.GetGame(id)
    if game.State != StateDraw {
        t.Errorf("expected draw, got %s", game.State)
    }
}

func TestInvalidMoves(t *testing.T) {
    svc := NewGameService()
    id := svc.CreateGame(3)
    if err := svc.MakeMove(id,0,0); err != ErrOutOfBounds {
        t.Errorf("expected ErrOutOfBounds, got %v", err)
    }
    svc.MakeMove(id,1,1)
    if err := svc.MakeMove(id,1,1); err != ErrCellOccupied {
        t.Errorf("expected ErrCellOccupied, got %v", err)
    }
    // simulate win then further move
    svc.MakeMove(id,1,2)
    svc.MakeMove(id,2,1)
    svc.MakeMove(id,2,2)
    svc.MakeMove(id,3,1) // X wins
    if err := svc.MakeMove(id,3,3); err != ErrGameEnded {
        t.Errorf("expected ErrGameEnded, got %v", err)
    }
}

func TestConcurrencyGameCreation(t *testing.T) {
    svc := NewGameService()
    var wg sync.WaitGroup
    ids := make(map[string]bool)
    mu := sync.Mutex{}

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            id := svc.CreateGame(3)
            mu.Lock()
            ids[id] = true
            mu.Unlock()
        }()
    }
    wg.Wait()
    if len(ids) != 100 {
        t.Errorf("expected 100 unique games, got %d", len(ids))
    }
}
