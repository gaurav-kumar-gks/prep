package data_structures

import (
	"fmt"
)

func DemonstrateStructs() {
	type Person struct {
		Name string
		Age  int
	}
	type User struct {
		Name     string `json:"name" xml:"name"`
		Age      int    `json:"age" xml:"age"`
		Password string `json:"-" xml:"-"`
	}	
	// p1 := Person{"John", 30}
	// p2 := Person{Name: "Jane", Age: 25}
	// p3 := new(Person) // Pointer to struct
	// p4 := &Person{Name: "Bob", Age: 35} // Pointer to struct

	// fmt.Println("Pointer Name:", (*p4).Name)
	// fmt.Println("Pointer Age:", p4.Age)

	
	type Node struct {
		Value int
		Next  *Node
	}
	
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2, Next: n1}
	
	fmt.Printf("Node 1: %v\n", n1)
	fmt.Printf("Node 2: %v\n", n2)
	
	type Test struct {
		Person
		Users []User
		Tasks chan string
		Data interface{}
		Process func(data interface{}) error
		Settings map[string]interface{}
	}
	
	t1 := Test{
		Person:   Person{Name: "Alice", Age: 28},
		Tasks: make(chan string),
		Data: "string data",
		Process: func(data interface{}) error {
			fmt.Printf("Processing: %v\n", data)
			return nil
		},
		Users: []User{
			{Name: "Admin1", Age: 30},
			{Name: "Admin2", Age: 35},
		},
		Settings: map[string]interface{}{
			"port":     8080,
		},
	}
	
	fmt.Printf("sample struct: %v\n", t1)

	p1 := Person{Name: "John", Age: 30}
	p2 := p1 // Creates a copy
	
	p2.Name = "Jane"
	fmt.Printf("Person 1: %v\n", p1)
	fmt.Printf("Person 2: %v\n", p2)
	
	type Point struct {
		X, Y int
	}
	
	pt1 := Point{1, 2}
	pt2 := Point{1, 2}
	pt3 := Point{2, 3}
	fmt.Printf("Points equal: %v\n", pt1 == pt2) // true
	fmt.Printf("Points equal: %v\n", pt1 == pt3) // false
}
