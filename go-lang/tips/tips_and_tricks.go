package tips

/*

1. STRING BUILDER
================

func buildString() string {
    res := "a"
    res += "b" // bad: creates new string each time
    var builder strings.Builder // good: pre-allocates memory
    for i := 0; i < 1000; i++ {
        builder.WriteString("a")
    }
    return builder.String()
}

2. SLICE CAPACITY
================

func processItems(items []int) []int {
    // result := make([]int, 0) // Bad: No capacity pre-allocation
    result := make([]int, 0, len(items)) //Pre-allocate capacity
    for _, item := range items {
        if item > 0 {
            result = append(result, item)
        }
    }
    return result
}


3. MAP INITIALIZATION
====================

func processMap(items []string) map[string]int {
    // result := make(map[string]int)  // No size pre-allocation    
    result := make(map[string]int, len(items)) // Pre-allocate size
    for _, item := range items {
        result[item]++
    }
    return result
}

*/