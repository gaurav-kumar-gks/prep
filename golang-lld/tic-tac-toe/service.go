package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

// GameService manages multiple games
type GameService struct {
    mu     sync.Mutex
    games  map[string]*Game
    nextID uint64
}

// NewGameService creates a new service instance
func NewGameService() *GameService {
    return &GameService{games: make(map[string]*Game)}
}

// CreateGame starts a game and returns its ID
func (s *GameService) CreateGame(size int) string {
    id := fmt.Sprintf("GAME-%d", atomic.AddUint64(&s.nextID, 1))
    game := NewGame(id, size)
    s.mu.Lock()
    s.games[id] = game
    s.mu.Unlock()
    return id
}

// MakeMove places the next symbol at (row,col); rows/cols are 1-based
func (s *GameService) MakeMove(gameID string, row, col int) error {
    s.mu.Lock()
    game, ok := s.games[gameID]
    s.mu.Unlock()
    if !ok {
        return ErrGameNotFound
    }

    game.mu.Lock()
    defer game.mu.Unlock()
    if game.State != StateRunning {
        return ErrGameEnded
    }
    if row < 1 || row > game.Size || col < 1 || col > game.Size {
        return ErrOutOfBounds
    }
    r, c := row-1, col-1
    if game.Board[r][c] != "" {
        return ErrCellOccupied
    }
    game.Board[r][c] = game.NextTurn
    // check win
    if checkWin(game.Board, game.NextTurn, r, c) {
        game.State = StateWon
        game.Winner = game.NextTurn
    } else if checkDraw(game.Board) {
        game.State = StateDraw
    } else {
        // swap turn
        if game.NextTurn == X {
            game.NextTurn = O
        } else {
            game.NextTurn = X
        }
    }
    return nil
}

// GetGame returns the current game state
func (s *GameService) GetGame(gameID string) (*Game, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    game, ok := s.games[gameID]
    if !ok {
        return nil, ErrGameNotFound
    }
    return game, nil
}

// checkDraw returns true if no empty cells remain
func checkDraw(board [][]PlayerSymbol) bool {
    for _, row := range board {
        for _, c := range row {
            if c == "" {
                return false
            }
        }
    }
    return true
}

// checkWin checks only rows, cols, diags through last move
func checkWin(board [][]PlayerSymbol, sym PlayerSymbol, r, c int) bool {
    n := len(board)
    win := true
    // row
    for j := 0; j < n; j++ {
        if board[r][j] != sym {
            win = false
            break
        }
    }
    if win { return true }
    // col
    win = true
    for i := 0; i < n; i++ {
        if board[i][c] != sym {
            win = false
            break
        }
    }
    if win { return true }
    // main diag
    if r == c {
        win = true
        for i := 0; i < n; i++ {
            if board[i][i] != sym {
                win = false
                break
            }
        }
        if win { return true }
    }
    // anti diag
    if r+c == n-1 {
        win = true
        for i := 0; i < n; i++ {
            if board[i][n-1-i] != sym {
                win = false
                break
            }
        }
        if win { return true }
    }
    return false
}