/*
LLD Design for Tic-Tac-Toe typically tests:
1) OOP modelling of game entities and turn-based logic
2) State management and win/draw detection algorithms
3) Thread-safety when managing multiple games concurrently

Typical Requirements:
- Initialize a 3x3 board
- Players alternate turns placing 'X' or 'O'
- Reject invalid moves: out-of-bounds, cell occupied, game already ended
- Determine and report win or draw
- Retrieve current game state and board

Typical Extensions & Solutions:
- Support variable board size N (generalize win check for N in a row)
- Add undo/redo via move history stack
- AI opponent using Minimax algorithm with pruning
- Persistence: save/load games from database
- Networked play: expose APIs with session management
*/

// models.go

// errors.go


// service.go


// main.go
package main

import (
    "fmt"
    "log"
)

func main() {
    svc := NewGameService()
    gameID := svc.CreateGame(3)
    fmt.Println("New game created:", gameID)

    moves := [][2]int{{1,1}, {1,2}, {2,2}, {1,3}, {3,3}}
    for _, mv := range moves {
        if err := svc.MakeMove(gameID, mv[0], mv[1]); err != nil {
            log.Fatalf("move failed: %v", err)
        }
        game, _ := svc.GetGame(gameID)
        fmt.Println(game)
    }

    // final state should be a win for X on main diag
}
