/*
LLD Design for Snake Game typically tests:
1) Stream processing and state management for continual moves
2) Data structure design for dynamic growth (deque/linked list) and collision detection
3) Algorithmic logic for food spawning, scoring, and game-over conditions

Typical Requirements:
- Initialize board with width, height, and list of food positions
- Move the snake in one of four directions ('U','D','L','R')
- On each move, return current score; -1 if game over
- Snake grows when it eats food; else moves forward by removing tail
- Game over if snake collides with wall or itself

Typical Extensions & Solutions:
- Random food generation when list is exhausted using uniform distribution
- Support multiple snakes (multiplayer) with collision detection between snakes
- Add obstacles on board as immovable blocks
- Add speed increase over time (dynamic tick interval)
- Expose API for real-time play with WebSocket streaming
*/
package main

import (
    "fmt"
)

func main() {
    food := []Position{{1,2},{0,1}}
    gs := NewGameService(3, 3, food)
    moves := []Direction{Right, Down, Right, Up, Left, Up}
    for _, mv := range moves {
        _, err := gs.Move(mv)
        if err != nil {
            fmt.Println("Move", mv, "-> Game Over")
            break
        }
        snake, sc, over := gs.GetState()
        fmt.Printf("Move %s: score=%d, snake=%v, over=%v\n", mv, sc, snake, over)
    }
}