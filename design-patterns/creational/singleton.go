package creational

import (
	"fmt"
	"sync"
)

// Singleton represents the singleton class
type Singleton struct {
	data string
}

var (
	instance *Singleton
	once     sync.Once
)

// GetInstance returns the singleton instance
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{data: "Initial data"}
	})
	return instance
}

// SetData sets the data
func (s *Singleton) SetData(data string) {
	s.data = data
}

// GetData gets the data
func (s *Singleton) GetData() string {
	return s.data
}

// SingletonDemo demonstrates the Singleton pattern
func SingletonDemo() {
	fmt.Println("--- Singleton Pattern Demo ---")
	
	// Get the singleton instance
	singleton1 := GetInstance()
	fmt.Printf("Initial data: %s\n", singleton1.GetData())
	
	// Modify the data
	singleton1.SetData("Modified data")
	fmt.Printf("Modified data through singleton1: %s\n", singleton1.GetData())
	
	// Get another reference to the singleton
	singleton2 := GetInstance()
	fmt.Printf("Data from singleton2: %s\n", singleton2.GetData())
	
	// Verify both references point to the same instance
	fmt.Printf("Are singleton1 and singleton2 the same instance? %v\n", singleton1 == singleton2)
} 