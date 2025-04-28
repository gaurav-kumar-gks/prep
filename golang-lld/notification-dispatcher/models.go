package main

// Notification represents a message to be sent
type Notification struct {
    ID        string
    Recipient string
    Message   string
    Attempts  int
}