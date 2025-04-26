package structural

import "fmt"

// Target defines the interface that the client expects
type Target interface {
	Request()
}

// Adaptee is the class that needs to be adapted
type Adaptee struct{}

// SpecificRequest is the method that needs to be adapted
func (a *Adaptee) SpecificRequest() {
	fmt.Println("Adaptee: specificRequest")
}

// Adapter adapts the Adaptee to the Target interface
type Adapter struct {
	adaptee *Adaptee
}

// NewAdapter creates a new Adapter
func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{adaptee: adaptee}
}

// Request implements the Target interface
func (a *Adapter) Request() {
	fmt.Println("Adapter: Converting Target request to Adaptee specificRequest")
	a.adaptee.SpecificRequest()
}

// Client uses the Target interface
type Client struct{}

// ClientCode demonstrates how the client uses the Target interface
func (c *Client) ClientCode(target Target) {
	fmt.Println("Client: I can work with Target interface")
	target.Request()
}

// AdapterDemo demonstrates the Adapter pattern
func AdapterDemo() {
	fmt.Println("Client: I can work with Target interface")
	
	// Create a concrete target
	target := &ConcreteTarget{}
	client := &Client{}
	client.ClientCode(target)
	
	fmt.Println("\nClient: I can't work with Adaptee directly")
	adaptee := &Adaptee{}
	fmt.Printf("Adaptee: %T\n", adaptee)
	
	fmt.Println("\nClient: But I can work with it via the Adapter")
	adapter := NewAdapter(adaptee)
	client.ClientCode(adapter)
}

// ConcreteTarget is a concrete implementation of the Target interface
type ConcreteTarget struct{}

// Request implements the Target interface
func (t *ConcreteTarget) Request() {
	fmt.Println("Target: request")
} 