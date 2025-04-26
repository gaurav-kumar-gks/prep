package data_structures

import (
	"fmt"
)

// DemonstratePointers shows pointer operations
func DemonstratePointers() {
	x := 42
	p := &x
	fmt.Printf("Value: %d, Address: %p\n", *p, p)
	pp := &p
	fmt.Printf("Value: %d, Address: %p, Pointer Address: %p\n", **pp, *pp, pp)
	
	np := new(int)
	*np = 100
	fmt.Printf("New pointer value: %d\n", *np)
	
	var nilPtr *int
	fmt.Printf("Nil pointer: %v\n", nilPtr)
	
	arr := [3]int{1, 2, 3}
	arrPtr := &arr[1]
	fmt.Printf("Array element pointer: %d\n", *arrPtr)
	
	slice := []int{1, 2, 3}
	slicePtr := &slice[1]
	fmt.Printf("Slice element pointer: %d\n", *slicePtr)
	
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "John", Age: 30}
	namePtr := &person.Name
	agePtr := &person.Age
	fmt.Printf("Person name pointer: %s\n", *namePtr)
	fmt.Printf("Person age pointer: %d\n", *agePtr)
	
	m := map[string]int{"a": 1}
	// Note: Can't take address of map value directly
	fmt.Printf("Map value: %d\n", m["a"])
	
	ch := make(chan int)
	chPtr := &ch
	fmt.Printf("Channel pointer: %v\n", chPtr)
	
	f := func() { fmt.Println("Function called") }
	fPtr := &f
	(*fPtr)()
	
	var i interface{} = 42
	iPtr := &i
	fmt.Printf("Interface pointer value: %v\n", *iPtr)
	
	const c = 42
	// Note: Can't take address of constant directly
	
	type MyInt int
	// Note: Can't take address of type directly
	
	// Note: Go doesn't support pointer arithmetic
	
	x1 := 42
	x2 := 42
	p1 := &x1
	p2 := &x2
	p3 := &x1
	fmt.Printf("Different pointers equal: %v\n", p1 == p2)
	fmt.Printf("Same pointers equal: %v\n", p1 == p3)
	
	ptrArr := [3]*int{&x1, &x2, nil}
	fmt.Printf("Array of pointers: %v\n", ptrArr)
	ptrSlice := []*int{&x1, &x2, nil}
	fmt.Printf("Slice of pointers: %v\n", ptrSlice)
	
	type PointerStruct struct {
		IntPtr    *int
		StringPtr *string
	}
	s := "hello"
	ps := PointerStruct{
		IntPtr:    &x1,
		StringPtr: &s,
	}
	fmt.Printf("Struct of pointers: %+v\n", ps)
	
	ptrMap := map[string]*int{
		"a": &x1,
		"b": &x2,
	}
	fmt.Printf("Map of pointers: %v\n", ptrMap)
	
	ptrCh := make(chan *int)
	fmt.Printf("Channel of pointers: %v\n", ptrCh)
}


func DemonstratePointerPatterns() {
	type Node struct {
		Value int
		Next  *Node
	}
	
	// Create linked list
	head := &Node{Value: 1}
	head.Next = &Node{Value: 2}
	head.Next.Next = &Node{Value: 3}
	
	// Traverse linked list
	current := head
	for current != nil {
		fmt.Printf("Node value: %d\n", current.Value)
		current = current.Next
	}
} 