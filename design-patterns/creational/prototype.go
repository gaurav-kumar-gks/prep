package creational

import (
	"fmt"
)

// Prototype defines the prototype interface
type Prototype interface {
	Clone() Prototype
	GetName() string
	SetName(name string)
}

// ConcretePrototype1 is a concrete prototype
type ConcretePrototype1 struct {
	name string
}

// Clone creates a copy of the prototype
func (p *ConcretePrototype1) Clone() Prototype {
	return &ConcretePrototype1{
		name: p.name,
	}
}

// GetName returns the name of the prototype
func (p *ConcretePrototype1) GetName() string {
	return p.name
}

// SetName sets the name of the prototype
func (p *ConcretePrototype1) SetName(name string) {
	p.name = name
}

// ConcretePrototype2 is a concrete prototype
type ConcretePrototype2 struct {
	name string
}

// Clone creates a copy of the prototype
func (p *ConcretePrototype2) Clone() Prototype {
	return &ConcretePrototype2{
		name: p.name,
	}
}

// GetName returns the name of the prototype
func (p *ConcretePrototype2) GetName() string {
	return p.name
}

// SetName sets the name of the prototype
func (p *ConcretePrototype2) SetName(name string) {
	p.name = name
}

// PrototypeRegistry manages prototypes
type PrototypeRegistry struct {
	prototypes map[string]Prototype
}

// NewPrototypeRegistry creates a new PrototypeRegistry
func NewPrototypeRegistry() *PrototypeRegistry {
	return &PrototypeRegistry{
		prototypes: make(map[string]Prototype),
	}
}

// Add adds a prototype to the registry
func (r *PrototypeRegistry) Add(id string, prototype Prototype) {
	r.prototypes[id] = prototype
}

// Get retrieves a prototype from the registry
func (r *PrototypeRegistry) Get(id string) Prototype {
	return r.prototypes[id]
}

// PrototypeDemo demonstrates the Prototype pattern
func PrototypeDemo() {
	// Create prototypes
	prototype1 := &ConcretePrototype1{}
	prototype1.SetName("Prototype 1")
	
	prototype2 := &ConcretePrototype2{}
	prototype2.SetName("Prototype 2")
	
	// Create registry
	registry := NewPrototypeRegistry()
	registry.Add("P1", prototype1)
	registry.Add("P2", prototype2)
	
	// Clone prototypes
	clone1 := registry.Get("P1").Clone()
	clone1.SetName("Clone 1")
	
	clone2 := registry.Get("P2").Clone()
	clone2.SetName("Clone 2")
	
	// Display results
	fmt.Printf("Original Prototype 1: %s\n", prototype1.GetName())
	fmt.Printf("Cloned Prototype 1: %s\n", clone1.GetName())
	fmt.Printf("Original Prototype 2: %s\n", prototype2.GetName())
	fmt.Printf("Cloned Prototype 2: %s\n", clone2.GetName())
} 