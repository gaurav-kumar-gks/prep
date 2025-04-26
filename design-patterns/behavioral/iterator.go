package behavioral

import "fmt"

// Iterator defines the iterator interface
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// Aggregate defines the aggregate interface
type Aggregate interface {
	CreateIterator() Iterator
}

// ConcreteIterator represents a concrete iterator
type ConcreteIterator struct {
	aggregate *ConcreteAggregate
	index     int
}

// NewConcreteIterator creates a new ConcreteIterator
func NewConcreteIterator(aggregate *ConcreteAggregate) *ConcreteIterator {
	return &ConcreteIterator{
		aggregate: aggregate,
		index:     0,
	}
}

// HasNext implements the Iterator interface
func (i *ConcreteIterator) HasNext() bool {
	return i.index < len(i.aggregate.items)
}

// Next implements the Iterator interface
func (i *ConcreteIterator) Next() interface{} {
	if i.HasNext() {
		item := i.aggregate.items[i.index]
		i.index++
		return item
	}
	return nil
}

// ConcreteAggregate represents a concrete aggregate
type ConcreteAggregate struct {
	items []string
}

// NewConcreteAggregate creates a new ConcreteAggregate
func NewConcreteAggregate() *ConcreteAggregate {
	return &ConcreteAggregate{
		items: make([]string, 0),
	}
}

// AddItem adds an item to the aggregate
func (a *ConcreteAggregate) AddItem(item string) {
	a.items = append(a.items, item)
}

// CreateIterator implements the Aggregate interface
func (a *ConcreteAggregate) CreateIterator() Iterator {
	return NewConcreteIterator(a)
}

// IteratorDemo demonstrates the Iterator pattern
func IteratorDemo() {
	aggregate := NewConcreteAggregate()
	
	aggregate.AddItem("Item 1")
	aggregate.AddItem("Item 2")
	aggregate.AddItem("Item 3")
	aggregate.AddItem("Item 4")
	aggregate.AddItem("Item 5")
	
	iterator := aggregate.CreateIterator()
	
	fmt.Println("--- Iterator Pattern Demo ---")
	fmt.Println("Iterating through items:")
	
	for iterator.HasNext() {
		item := iterator.Next()
		fmt.Printf("Item: %v\n", item)
	}
} 