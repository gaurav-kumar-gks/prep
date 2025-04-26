package behavioral

import "fmt"

// CaffeineBeverage defines the abstract class
type CaffeineBeverage interface {
	Prepare()
	BoilWater()
	PourInCup()
	Brew()
	AddCondiments()
	CustomerWantsCondiments() bool
}

// BaseBeverage provides default implementations
type BaseBeverage struct{}

// BoilWater provides default implementation
func (b *BaseBeverage) BoilWater() {
	fmt.Println("Boiling water")
}

// PourInCup provides default implementation
func (b *BaseBeverage) PourInCup() {
	fmt.Println("Pouring into cup")
}

// CustomerWantsCondiments provides default implementation
func (b *BaseBeverage) CustomerWantsCondiments() bool {
	return true
}

// Prepare defines the template method
func Prepare(beverage CaffeineBeverage) {
	beverage.BoilWater()
	beverage.Brew()
	beverage.PourInCup()
	if beverage.CustomerWantsCondiments() {
		beverage.AddCondiments()
	}
}

// Coffee represents a concrete class
type Coffee struct {
	BaseBeverage
}

// Brew implements the CaffeineBeverage interface
func (c *Coffee) Brew() {
	fmt.Println("Dripping coffee through filter")
}

// AddCondiments implements the CaffeineBeverage interface
func (c *Coffee) AddCondiments() {
	fmt.Println("Adding sugar and milk")
}

// Tea represents a concrete class
type Tea struct {
	BaseBeverage
}

// Brew implements the CaffeineBeverage interface
func (t *Tea) Brew() {
	fmt.Println("Steeping the tea")
}

// AddCondiments implements the CaffeineBeverage interface
func (t *Tea) AddCondiments() {
	fmt.Println("Adding lemon")
}

// TemplateMethodDemo demonstrates the Template Method pattern
func TemplateMethodDemo() {
	fmt.Println("--- Template Method Pattern Demo ---")
	
	fmt.Println("\nMaking coffee:")
	coffee := &Coffee{}
	Prepare(coffee)
	
	fmt.Println("\nMaking tea:")
	tea := &Tea{}
	Prepare(tea)
} 