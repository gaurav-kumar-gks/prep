package main

import (
	"fmt"
	"design-patterns/structural"
)

func main() {
	fmt.Println("=== Structural Design Patterns ===")
	
	fmt.Println("\n--- Adapter Pattern ---")
	structural.AdapterDemo()
	
	fmt.Println("\n--- Bridge Pattern ---")
	structural.BridgeDemo()
	
	fmt.Println("\n--- Composite Pattern ---")
	structural.CompositeDemo()
	
	fmt.Println("\n--- Decorator Pattern ---")
	structural.DecoratorDemo()
	
	fmt.Println("\n--- Facade Pattern ---")
	structural.FacadeDemo()
	
	fmt.Println("\n--- Proxy Pattern ---")
	structural.ProxyDemo()
} 