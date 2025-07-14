package stdlib


import (
    "encoding/json"
    "fmt"
    "os"
    "strings"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// JSON MARSHALING AND UNMARSHALING
// ===============

func main() {
    person := Person{
        Name: "John Doe",
        Age:  30,
    }
    
    // Marshal to JSON
    jsonData, err := json.Marshal(person)
    // jsonPretty, err := json.MarshalIndent(person, "", "  ")
    if err != nil {
        fmt.Println("Error marshaling:", err)
        return
    }
    fmt.Println(string(jsonData)) // {"name":"John Doe","age":30}
    
    // Unmarshal from JSON
    jsonStr := `{"name":"Jane Doe","age":25}`
    var person1 Person
    unmarshalerr := json.Unmarshal([]byte(jsonStr), &person1)
    if unmarshalerr != nil {
        fmt.Println("Error unmarshaling:", unmarshalerr)
        return
    }
}

// JSON FILE HANDLING
// ===============

func LoadPersonFromFile(filename string) (*Person, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("error opening file: %w", err)
    }
    defer file.Close()
    decoder := json.NewDecoder(file)
    var person Person
    if err := decoder.Decode(&person); err != nil {
        return nil, fmt.Errorf("error decoding JSON: %w", err)
    }
    return &person, nil
}

func SavePersonToFile(person *Person, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return fmt.Errorf("error creating file: %w", err)
    }
    defer file.Close()
    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    if err := encoder.Encode(person); err != nil {
        return fmt.Errorf("error encoding JSON: %w", err)
    }
    return nil
}


// JSON STREAMING
// ===============
// Using json.Decoder and json.Encoder for streaming JSON:
func JsonStreaming() {
    jsonStr := `{"name":"John","age":30}{"name":"Jane","age":25}{"name":"Bob","age":40}`
    
    reader := strings.NewReader(jsonStr)
    decoder := json.NewDecoder(reader)
    
    for {
        var person Person
        err := decoder.Decode(&person)
        if err != nil {
            // End of input
            break
        }
        fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
    }
    
    var buf strings.Builder
    encoder := json.NewEncoder(&buf)
    people := []Person{
        {Name: "Alice", Age: 35},
        {Name: "Charlie", Age: 45},
    }
    
    for _, person := range people {
        err := encoder.Encode(person)
        if err != nil {
            fmt.Println("Error encoding:", err)
            return
        }
    }
    
    fmt.Println(buf.String())
    // {"name":"Alice","age":35}
    // {"name":"Charlie","age":45}
}

