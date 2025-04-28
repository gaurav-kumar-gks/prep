package main

import (
    "fmt"
    "sync"
)

// PlayerSymbol represents 'X' or 'O'
type PlayerSymbol string

const (
    X PlayerSymbol = "X"
    O PlayerSymbol = "O"
)

// GameState indicates the status of a game
type GameState string

const (
    StateRunning GameState = "RUNNING"
    StateDraw    GameState = "DRAW"
    StateWon     GameState = "WON"
)

// Game holds board and status
type Game struct {
    ID        string
    Size      int
    Board     [][]PlayerSymbol
    NextTurn  PlayerSymbol
    State     GameState
    Winner    PlayerSymbol
    mu        sync.Mutex
}

// NewGame initializes a new game of given size
func NewGame(id string, size int) *Game {
    board := make([][]PlayerSymbol, size)
    for i := range board {
        board[i] = make([]PlayerSymbol, size)
    }
    return &Game{
        ID:       id,
        Size:     size,
        Board:    board,
        NextTurn: X,
        State:    StateRunning,
    }
}

func (g *Game) String() string {
    s := fmt.Sprintf("Game %s (Next: %s, State: %s)\n", g.ID, g.NextTurn, g.State)
    for _, row := range g.Board {
        for _, c := range row {
            cell := string(c)
            if c == "" {
                cell = "."
            }
            s += cell + " "
        }
        s += "\n"
    }
    return s
}
