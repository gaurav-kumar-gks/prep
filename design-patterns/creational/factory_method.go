package creational

import "fmt"

// Product defines the interface for products
type Product interface {
	Operation() string
}

// ConcreteProduct1 is a concrete product
type ConcreteProduct1 struct{}

// Operation implements the Product interface
func (p *ConcreteProduct1) Operation() string {
	return "{Result of the ConcreteProduct1}"
}

// ConcreteProduct2 is a concrete product
type ConcreteProduct2 struct{}

// Operation implements the Product interface
func (p *ConcreteProduct2) Operation() string {
	return "{Result of the ConcreteProduct2}"
}

// Creator defines the creator interface
type Creator interface {
	FactoryMethod() Product
	SomeOperation() string
}

// ConcreteCreator1 is a concrete creator
type ConcreteCreator1 struct{}

// FactoryMethod creates a ConcreteProduct1
func (c *ConcreteCreator1) FactoryMethod() Product {
	return &ConcreteProduct1{}
}

// SomeOperation performs some operation
func (c *ConcreteCreator1) SomeOperation() string {
	product := c.FactoryMethod()
	return "Creator: The same creator's code has just worked with " + product.Operation()
}

// ConcreteCreator2 is a concrete creator
type ConcreteCreator2 struct{}

// FactoryMethod creates a ConcreteProduct2
func (c *ConcreteCreator2) FactoryMethod() Product {
	return &ConcreteProduct2{}
}

// SomeOperation performs some operation
func (c *ConcreteCreator2) SomeOperation() string {
	product := c.FactoryMethod()
	return "Creator: The same creator's code has just worked with " + product.Operation()
}

// ClientCode demonstrates how the client uses the creator
func ClientCode(creator Creator) {
	fmt.Printf("Client: I'm not aware of the creator's class, but it still works.\n%s\n", creator.SomeOperation())
}

// FactoryMethodDemo demonstrates the Factory Method pattern
func FactoryMethodDemo() {
	fmt.Println("App: Launched with the ConcreteCreator1.")
	ClientCode(&ConcreteCreator1{})
	fmt.Println()

	fmt.Println("App: Launched with the ConcreteCreator2.")
	ClientCode(&ConcreteCreator2{})
} 