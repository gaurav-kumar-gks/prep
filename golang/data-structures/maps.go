package data_structures

import (
	"fmt"
)

/*

MAP METHODS
==========
1. make() - Create map
   m := make(map[string]int)

2. len() - Get map length
   len(m)

3. delete() - Delete key
   delete(m, "key")

4. [key] - Access value
   value := m["key"]

5. [key] = value - Set value
   m["key"] = value

6. value, ok = m[key] - Check existence
   value, ok := m["key"]

7. range - Iterate map
   for k, v := range m

8. clear() - Clear map (Go 1.21+)
   clear(m)

9. maps.Clone() - Clone map (Go 1.21+)
   clone := maps.Clone(m)

10. maps.Copy() - Copy map (Go 1.21+)
    maps.Copy(dst, src)

11. maps.Delete() - Delete key (Go 1.21+)
    maps.Delete(m, "key")

12. maps.Equal() - Compare maps (Go 1.21+)
    maps.Equal(m1, m2)

13. maps.Keys() - Get keys (Go 1.21+)
    keys := maps.Keys(m)

14. maps.Values() - Get values (Go 1.21+)
    values := maps.Values(m)
	for v := range values {
		fmt.Println(v)
	}
	fmt.Println("Values: ", slices.Sorted(values))	

15. maps.EqualFunc() - Compare with function (Go 1.21+)
    maps.EqualFunc(m1, m2, func(v1, v2 int) bool { return v1 == v2 })

16. maps.CloneFunc() - Clone with function (Go 1.21+)
    maps.CloneFunc(m, func(v int) int { return v * 2 })

17. maps.CopyFunc() - Copy with function (Go 1.21+)
    maps.CopyFunc(dst, src, func(v int) int { return v * 2 })

18. maps.DeleteFunc() - Delete with function (Go 1.21+)
    maps.DeleteFunc(m, func(k string, v int) bool { return v > 10 })

19. maps.KeysFunc() - Get keys with function (Go 1.21+)
    maps.KeysFunc(m, func(k string, v int) bool { return v > 10 })

20. maps.ValuesFunc() - Get values with function (Go 1.21+)
    maps.ValuesFunc(m, func(k string, v int) bool { return v > 10 })
*/

// DemonstrateMaps shows map operations
func DemonstrateMaps() {
	fmt.Println("\n=== Maps ===")
	
	// make() creates a map with zero length and capacity, used for dynamic maps, non string keys, nil values
	// Literal creates a map with initial values, used for static maps, non nil values
	m1 := make(map[string]int)
	m2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	
	m1["hello"] = 42
	m1["world"] = 100
	if value, ok := m1["hello"]; ok {
		fmt.Printf("Found value: %d\n", value)
	}
	
	// for k, v := range m2 {
	// 	 fmt.Printf("%s: %d\n", k, v)
	// }
	
	// delete(m1, "hello")
	// len(m1)
	
	// Map copying
	m3 := make(map[string]int)
	for k, v := range m2 {
		m3[k] = v
	}
	
	// Map comparison
	equal := true
	if len(m2) != len(m3) {
		equal = false
	} else {
		for k, v := range m2 {
			if m3[k] != v {
				equal = false
				break
			}
		}
	}
	fmt.Printf("Maps equal: %v\n", equal)
	// alternative: use maps.Equal() in Go 1.21+
	// equal := maps.Equal(m2, m3)
	// fmt.Printf("Maps equal: %v\n", equal)


	// Map can have struct / slices / map / interface / function / channel / pointer values 
	// Map can have custom keys
	type Key struct {
		ID   int
		Name string
	}
	customKeys := make(map[Key]string)
	customKeys[Key{1, "one"}] = "first"
	
	// new keyword: 
	// new allocates memory for a value of the specified type and returns a pointer to it
	// new is used for creating pointers to values
	// e.g. new(int) creates a pointer to an int
	// e.g. new([]int) creates a pointer to a slice of int
	// e.g. new(map[string]int) creates a pointer to a map of string to int
	// e.g. new(chan int) creates a pointer to a channel of int
	anyValues := map[string]interface{}{
		"string": "hello",
		"int":    42,
		"slice":  []int{1, 2, 3},
		"map":    map[string]int{"a": 1},
	}
	fmt.Printf("Any values: %v\n", anyValues)
}
