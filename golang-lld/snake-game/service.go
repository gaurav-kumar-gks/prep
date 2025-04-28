package main

import (
	"sync"
)

// GameService encapsulates snake game logic
type GameService struct {
	width     int
	height    int
	food      []Position
	foodIndex int
	score     int
	snake     *Deque           // Using Deque for efficient snake movement
	occupied  map[Position]bool // quick lookup of snake body
	mu        sync.Mutex
	gameOver  bool
}

// NewGameService initializes the game with board size and food list
func NewGameService(width, height int, food []Position) *GameService {
	initial := NewDeque(10) // Start with capacity 10
	initial.PushFront(Position{0, 0})
	occ := map[Position]bool{{0, 0}: true}
	return &GameService{
		width:    width,
		height:   height,
		food:     food,
		snake:    initial,
		occupied: occ,
	}
}

// Move processes a direction and returns updated score or -1 if game over
func (gs *GameService) Move(dir Direction) (int, error) {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	if gs.gameOver {
		return -1, ErrGameOver
	}
	
	head := gs.snake.Front()
	var newHead Position
	switch dir {
	case Up:
		newHead = Position{head.Row - 1, head.Col}
	case Down:
		newHead = Position{head.Row + 1, head.Col}
	case Left:
		newHead = Position{head.Row, head.Col - 1}
	case Right:
		newHead = Position{head.Row, head.Col + 1}
	default:
		return gs.score, nil
	}
	
	// check wall collision
	if newHead.Row < 0 || newHead.Row >= gs.height || newHead.Col < 0 || newHead.Col >= gs.width {
		gs.gameOver = true
		return -1, ErrGameOver
	}
	
	// check self collision
	if gs.occupied[newHead] {
		gs.gameOver = true
		return -1, ErrGameOver
	}
	
	// add new head
	gs.snake.PushFront(newHead)
	gs.occupied[newHead] = true
	
	// check food
	ateFood := false
	if gs.foodIndex < len(gs.food) && 
	   newHead.Row == gs.food[gs.foodIndex].Row && 
	   newHead.Col == gs.food[gs.foodIndex].Col {
		gs.score++
		gs.foodIndex++
		ateFood = true
	}
	
	// If didn't eat food, remove tail
	if !ateFood {
		tail := gs.snake.PopBack()
		delete(gs.occupied, tail)
	}
	
	return gs.score, nil
}

// GetState returns current snake and score
func (gs *GameService) GetState() ([]Position, int, bool) {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	return gs.snake.ToSlice(), gs.score, gs.gameOver
}