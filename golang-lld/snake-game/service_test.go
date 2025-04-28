package main

import (
    "testing"
)

func TestSimpleMoves(t *testing.T) {
    food := []Position{{2,0}}
    gs := NewGameService(3,3,food)
    tests := []struct{dir Direction; expectScore int; expectOver bool}{
        {Right, 0, false},
        {Right, 0, false},
        {Down, 1, false}, // eats food
        {Left, 1, false},
        {Up, -1, true},   // collision with self
    }

    for _, tc := range tests {
        score, err := gs.Move(tc.dir)
        if tc.expectOver {
            if err == nil {
                t.Errorf("expected game over on %v", tc.dir)
            }
        } else {
            if score != tc.expectScore {
                t.Errorf("dir %v expected score %d, got %d", tc.dir, tc.expectScore, score)
            }
        }
    }
}

func TestWallCollision(t *testing.T) {
    gs := NewGameService(2,2,nil)
    _, err := gs.Move(Left)
    if err == nil {
        t.Errorf("expected wall collision")
    }
}

func TestGrowth(t *testing.T) {
    food := []Position{{0,1},{0,2}}
    gs := NewGameService(3,1,food)
    sc,_ := gs.Move(Right) // eat1, len=2
    if sc!=1 { t.Errorf("expected score1, got %d", sc)}
    sc,_ = gs.Move(Right) // eat2 len=3
    if sc!=2 { t.Errorf("expected score2, got %d", sc)}
    // next Right -> collision wall\ n    if _,err:=gs.Move(Right); err==nil { t.Errorf("expected collision") }
}
