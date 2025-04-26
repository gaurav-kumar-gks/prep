package data_structures

import (
	"fmt"
	"slices"
	"sort"
	"unsafe"
)


func DemonstrateArrays() {
	// var arr1 [5]int
	arr2 := [5]int{1, 2, 3, 4, 5}
	// arr3 := [...]int{1, 2, 3} // Size inferred
	// Array size is part of type
	// var arr4 [4]int = arr2 // Error: different sizes
	arr4 := arr2 // Creates a copy
	arr4[0] = 10
	arr5 := arr2[:] // Creates a slice from array
	fmt.Println("length of arr2:", len(arr5))
	fmt.Println("capacity of arr2:", cap(arr2))
	if arr2 == arr4 {
		fmt.Println("arr2 and arr4 are equal")
	} else {
		fmt.Println("arr2 and arr4 are not equal")
	}
	fmt.Printf("Array size: %v bytes\n", unsafe.Sizeof(arr2))
	fmt.Printf("Array element size: %v bytes\n", unsafe.Sizeof(arr2[0]))
	fmt.Println("Array iteration:")
	for i, v := range arr2 {
		fmt.Printf("arr2[%v] = %v\n", i, v)
	}
}

// DemonstrateSlices shows slice concepts
func DemonstrateSlices() {
	slice := []int{1, 2, 3, 4, 5}
	var slice1 []int
	slice3 := make([]int, 5, 10)
	slice4 := slice[1:3]
	// make([]T, len, cap) creates a slice of type T
	// make not only allocates memory but also initializes it
	// why make: because it imporves performance by preallocating capacity for slices 
	// and thus reducing reallocations
	// used for slices, maps, and channels (composite types)
	println("length of slice1:", len(slice1))
	println("length of slice2:", len(slice3))
	println("capacity of slice2:", cap(slice4))
	
	arr := [5]int{1, 2, 3, 4, 5}
	slice5 := arr[1:3]
	
	slice6 := slice5[:cap(slice5)]
	fmt.Printf("Slice 6: %v, len=%v, cap=%v\n", slice6, len(slice6), cap(slice6))
	

	slice = append(slice, 6, 7, 8)
	otherSlice := []int{9, 10}
	slice = append(slice, otherSlice...)

	// slice7 := []int{1, 2, 3, 4, 5}
	// slice8 := slice7[1:3] // creates a new slice
	// slice8 shares the same underlying array as slice7
	// modifying slice8 will affect slice7 and vice versa

	copiedSlice := make([]int, len(slice))
	copy(copiedSlice, slice)
	copiedSlice[0] = 10 // does not affect original slice
	
	// For small slices, the difference is negligible
	// For medium slices, copy is generally faster than append
	// For very large slices, unsafe.Copy can be faster, but it's unsafe and not recommended for general use
	// For structs without pointers, a simple copy is sufficient; deep copy is only needed for structs with pointers

	unsorted := []int{5, 2, 8, 1, 9, 3}
	sort.Ints(unsorted)
	sort.Slice(unsorted, func(i, j int) bool {
		return unsorted[i] > unsorted[j]
	})
	
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := sort.SearchInts(sorted, 5)
	fmt.Printf("Found 5 at index: %d\n", index)
	index = sort.Search(len(sorted), func(i int) bool {
		return sorted[i] >= 5
	})
	fmt.Printf("Custom search found 5 at index: %d\n", index)
	
	clear(slice)
	slice = []int{1, 2, 3, 4, 5}
	slice = slices.Delete(slice, 1, 3) // delete elements 1 to 3
	
	slice = []int{1, 2, 3, 4, 5}
	slice = slices.Replace(slice, 1, 3, 6, 7) // replace elements 1 to 3 with 6 and 7
}

/*

Difference between Array and Slice and when to use both, including performance considerations:

1. Array:
   - Fixed size, defined at compile time
   - Value type, copied when passed to functions
   - Memory layout: contiguous block of memory
   - Allocated on the stack
   - Passed by value
   - Size is part of the type
   - Use when size is known and fixed
   - Performance: faster for small arrays, less overhead
2. Slice:
   - Dynamic size, defined at runtime
   - Reference type, pointer to an underlying array
   - Allocated on the heap
   - Passde by reference
   - Memory layout: slice header (pointer, length, capacity)
   - Size is not part of the type
   - Use when size is unknown or variable
   - Performance: slower for small slices, more overhead

*/