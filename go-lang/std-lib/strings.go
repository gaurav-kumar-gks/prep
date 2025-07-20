package stdlib

/*

import "strings"

func main() {
    str := "Hello, World!"
    upper := strings.ToUpper(str)    // "HELLO, WORLD!"
    lower := strings.ToLower(str)    // "hello, world!"
    title := strings.Title(str)      // "Hello, World!"
    
    trimmed := strings.TrimSpace("  hello  ")  // "hello"
    left := strings.TrimLeft("  hello  ", " ") // "hello  "
    right := strings.TrimRight("  hello  ", " ") // "  hello"
    
    padded := strings.Repeat("0", 5)  // "00000"
    
    parts := strings.Split("a,b,c", ",")  // []string{"a", "b", "c"}
    parts = strings.Fields("  hello  world  ")  // []string{"hello", "world"}
    joined := strings.Join(parts, "-")     // "a-b-c"
    

    hasHello := strings.Contains(str, "Hello")  // true
    hasGolang := strings.Contains(str, "golang") // false
    contains := strings.Contains(strings.ToLower(str), "world") // true
    
    count := strings.Count("hello hello", "hello") // 2
    index := strings.Index(str, "World")  // 7
    lastIndex := strings.LastIndex(str, "l") // 10
    
    hasPrefix := strings.HasPrefix(str, "Hello") // true
    hasSuffix := strings.HasSuffix(str, "!")     // true

    equal := "hello" == "hello"  // true
    equalIgnoreCase := strings.EqualFold("Hello", "hello") // true
    result := strings.Compare("a", "b")  // -1
    result = strings.Compare("b", "a")   // 1
    result = strings.Compare("a", "a")   // 0

    str := "hello world world"
    replaced := strings.Replace(str, "world", "golang", 1) // "hello golang world"
    allReplaced := strings.ReplaceAll(str, "world", "golang") // "hello golang golang"
    replacer := strings.NewReplacer(
        "hello", "hi",
        "world", "golang",
    )
    replaced = replacer.Replace(str)  // "hi golang golang"
}

*/