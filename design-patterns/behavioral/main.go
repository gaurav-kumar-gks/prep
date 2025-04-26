package main

import (
	"fmt"
	"design-patterns/behavioral"
)

func main() {
	fmt.Println("=== Behavioral Design Patterns ===")
	
	fmt.Println("\n--- Command Pattern ---")
	behavioral.CommandDemo()
	
	fmt.Println("\n--- Observer Pattern ---")
	behavioral.ObserverDemo()
	
	fmt.Println("\n--- State Pattern ---")
	behavioral.StateDemo()
	
	fmt.Println("\n--- Template Method Pattern ---")
	behavioral.TemplateMethodDemo()
	
	fmt.Println("\n--- Chain of Responsibility Pattern ---")
	behavioral.ChainOfResponsibilityDemo()
	
	fmt.Println("\n--- Mediator Pattern ---")
	behavioral.MediatorDemo()
	
	fmt.Println("\n--- Visitor Pattern ---")
	behavioral.VisitorDemo()
	
	fmt.Println("\n--- Iterator Pattern ---")
	behavioral.IteratorDemo()
} 