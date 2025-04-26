package behavioral

import "fmt"

// Element defines the element interface
type Element interface {
	Accept(visitor Visitor)
}

// Visitor defines the visitor interface
type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

// ConcreteElementA represents a concrete element
type ConcreteElementA struct{}

// Accept implements the Element interface
func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

// OperationA performs operation A
func (e *ConcreteElementA) OperationA() string {
	return "ConcreteElementA"
}

// ConcreteElementB represents a concrete element
type ConcreteElementB struct{}

// Accept implements the Element interface
func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

// OperationB performs operation B
func (e *ConcreteElementB) OperationB() string {
	return "ConcreteElementB"
}

// ConcreteVisitor1 represents a concrete visitor
type ConcreteVisitor1 struct{}

// VisitConcreteElementA implements the Visitor interface
func (v *ConcreteVisitor1) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Printf("ConcreteVisitor1: Visiting %s\n", element.OperationA())
}

// VisitConcreteElementB implements the Visitor interface
func (v *ConcreteVisitor1) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Printf("ConcreteVisitor1: Visiting %s\n", element.OperationB())
}

// ConcreteVisitor2 represents a concrete visitor
type ConcreteVisitor2 struct{}

// VisitConcreteElementA implements the Visitor interface
func (v *ConcreteVisitor2) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Printf("ConcreteVisitor2: Visiting %s\n", element.OperationA())
}

// VisitConcreteElementB implements the Visitor interface
func (v *ConcreteVisitor2) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Printf("ConcreteVisitor2: Visiting %s\n", element.OperationB())
}

// ObjectStructure represents the object structure
type ObjectStructure struct {
	elements []Element
}

// NewObjectStructure creates a new ObjectStructure
func NewObjectStructure() *ObjectStructure {
	return &ObjectStructure{
		elements: make([]Element, 0),
	}
}

// Attach adds an element to the structure
func (o *ObjectStructure) Attach(element Element) {
	o.elements = append(o.elements, element)
}

// Detach removes an element from the structure
func (o *ObjectStructure) Detach(element Element) {
	for i, e := range o.elements {
		if e == element {
			o.elements = append(o.elements[:i], o.elements[i+1:]...)
			break
		}
	}
}

// Accept accepts a visitor
func (o *ObjectStructure) Accept(visitor Visitor) {
	for _, element := range o.elements {
		element.Accept(visitor)
	}
}

// VisitorDemo demonstrates the Visitor pattern
func VisitorDemo() {
	objectStructure := NewObjectStructure()
	
	objectStructure.Attach(&ConcreteElementA{})
	objectStructure.Attach(&ConcreteElementB{})
	
	visitor1 := &ConcreteVisitor1{}
	visitor2 := &ConcreteVisitor2{}
	
	fmt.Println("--- Visitor Pattern Demo ---")
	
	fmt.Println("\nVisiting with ConcreteVisitor1:")
	objectStructure.Accept(visitor1)
	
	fmt.Println("\nVisiting with ConcreteVisitor2:")
	objectStructure.Accept(visitor2)
} 