package main

// Direction represents the direction of snake movement
type Direction int

const (
    Up Direction = iota
    Down
    Left
    Right
)

// Position represents a point on the game board
type Position struct {
    Row, Col int
}

// Deque implements a double-ended queue for efficient snake movement
type Deque struct {
    data     []Position
    head     int
    tail     int
    size     int
    capacity int
}

// NewDeque creates a new deque with initial capacity
func NewDeque(capacity int) *Deque {
    return &Deque{
        data:     make([]Position, capacity),
        capacity: capacity,
    }
}

// PushFront adds an element to the front of the deque
func (d *Deque) PushFront(pos Position) {
    if d.size == d.capacity {
        // Grow the deque
        newData := make([]Position, d.capacity*2)
        // Copy elements maintaining order
        for i := 0; i < d.size; i++ {
            newData[i] = d.data[(d.head+i)%d.capacity]
        }
        d.data = newData
        d.head = 0
        d.tail = d.size - 1
        d.capacity *= 2
    }
    
    d.head = (d.head - 1 + d.capacity) % d.capacity
    d.data[d.head] = pos
    d.size++
}

// PopBack removes and returns the last element
func (d *Deque) PopBack() Position {
    if d.size == 0 {
        panic("deque is empty")
    }
    pos := d.data[d.tail]
    d.tail = (d.tail - 1 + d.capacity) % d.capacity
    d.size--
    return pos
}

// Front returns the first element without removing it
func (d *Deque) Front() Position {
    if d.size == 0 {
        panic("deque is empty")
    }
    return d.data[d.head]
}

// Back returns the last element without removing it
func (d *Deque) Back() Position {
    if d.size == 0 {
        panic("deque is empty")
    }
    return d.data[d.tail]
}

// ToSlice returns all elements in order
func (d *Deque) ToSlice() []Position {
    result := make([]Position, d.size)
    for i := 0; i < d.size; i++ {
        result[i] = d.data[(d.head+i)%d.capacity]
    }
    return result
}

// Size returns the number of elements
func (d *Deque) Size() int {
    return d.size
}
