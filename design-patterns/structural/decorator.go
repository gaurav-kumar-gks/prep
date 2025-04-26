package structural

import "fmt"

// Coffee defines the component interface
type Coffee interface {
	GetCost() float64
	GetDescription() string
}

// SimpleCoffee is a concrete component
type SimpleCoffee struct{}

// GetCost implements the Coffee interface
func (c *SimpleCoffee) GetCost() float64 {
	return 1.0
}

// GetDescription implements the Coffee interface
func (c *SimpleCoffee) GetDescription() string {
	return "Simple coffee"
}

// CoffeeDecorator is the abstract decorator
type CoffeeDecorator struct {
	decoratedCoffee Coffee
}

// NewCoffeeDecorator creates a new CoffeeDecorator
func NewCoffeeDecorator(coffee Coffee) *CoffeeDecorator {
	return &CoffeeDecorator{decoratedCoffee: coffee}
}

// GetCost implements the Coffee interface
func (d *CoffeeDecorator) GetCost() float64 {
	return d.decoratedCoffee.GetCost()
}

// GetDescription implements the Coffee interface
func (d *CoffeeDecorator) GetDescription() string {
	return d.decoratedCoffee.GetDescription()
}

// MilkDecorator is a concrete decorator
type MilkDecorator struct {
	*CoffeeDecorator
}

// NewMilkDecorator creates a new MilkDecorator
func NewMilkDecorator(coffee Coffee) *MilkDecorator {
	return &MilkDecorator{
		CoffeeDecorator: NewCoffeeDecorator(coffee),
	}
}

// GetCost implements the Coffee interface
func (d *MilkDecorator) GetCost() float64 {
	return d.decoratedCoffee.GetCost() + 0.5
}

// GetDescription implements the Coffee interface
func (d *MilkDecorator) GetDescription() string {
	return d.decoratedCoffee.GetDescription() + ", with milk"
}

// SugarDecorator is a concrete decorator
type SugarDecorator struct {
	*CoffeeDecorator
}

// NewSugarDecorator creates a new SugarDecorator
func NewSugarDecorator(coffee Coffee) *SugarDecorator {
	return &SugarDecorator{
		CoffeeDecorator: NewCoffeeDecorator(coffee),
	}
}

// GetCost implements the Coffee interface
func (d *SugarDecorator) GetCost() float64 {
	return d.decoratedCoffee.GetCost() + 0.2
}

// GetDescription implements the Coffee interface
func (d *SugarDecorator) GetDescription() string {
	return d.decoratedCoffee.GetDescription() + ", with sugar"
}

// WhippedCreamDecorator is a concrete decorator
type WhippedCreamDecorator struct {
	*CoffeeDecorator
}

// NewWhippedCreamDecorator creates a new WhippedCreamDecorator
func NewWhippedCreamDecorator(coffee Coffee) *WhippedCreamDecorator {
	return &WhippedCreamDecorator{
		CoffeeDecorator: NewCoffeeDecorator(coffee),
	}
}

// GetCost implements the Coffee interface
func (d *WhippedCreamDecorator) GetCost() float64 {
	return d.decoratedCoffee.GetCost() + 0.7
}

// GetDescription implements the Coffee interface
func (d *WhippedCreamDecorator) GetDescription() string {
	return d.decoratedCoffee.GetDescription() + ", with whipped cream"
}

// DecoratorDemo demonstrates the Decorator pattern
func DecoratorDemo() {
	// Create a simple coffee
	coffee := &SimpleCoffee{}
	fmt.Printf("Cost: $%.2f; Description: %s\n", coffee.GetCost(), coffee.GetDescription())

	// Add milk
	coffee = NewMilkDecorator(coffee)
	fmt.Printf("Cost: $%.2f; Description: %s\n", coffee.GetCost(), coffee.GetDescription())

	// Add sugar
	coffee = NewSugarDecorator(coffee)
	fmt.Printf("Cost: $%.2f; Description: %s\n", coffee.GetCost(), coffee.GetDescription())

	// Add whipped cream
	coffee = NewWhippedCreamDecorator(coffee)
	fmt.Printf("Cost: $%.2f; Description: %s\n", coffee.GetCost(), coffee.GetDescription())
} 