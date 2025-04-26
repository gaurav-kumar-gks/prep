package main

import (
	"fmt"
	"design-patterns/creational"
)

func main() {
	fmt.Println("=== Creational Design Patterns ===")
	
	fmt.Println("\n--- Abstract Factory Pattern ---")
	creational.AbstractFactoryDemo()
	
	fmt.Println("\n--- Builder Pattern ---")
	creational.BuilderDemo()
	
	fmt.Println("\n--- Factory Method Pattern ---")
	creational.FactoryMethodDemo()
	
	fmt.Println("\n--- Prototype Pattern ---")
	creational.PrototypeDemo()
	
	fmt.Println("\n--- Singleton Pattern ---")
	creational.SingletonDemo()
} 