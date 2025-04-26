package behavioral

import "fmt"

// Handler defines the handler interface
type Handler interface {
	SetNext(handler Handler)
	Handle(request string) string
}

// BaseHandler provides default implementation
type BaseHandler struct {
	next Handler
}

// SetNext sets the next handler
func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

// Handle provides default implementation
func (h *BaseHandler) Handle(request string) string {
	if h.next != nil {
		return h.next.Handle(request)
	}
	return ""
}

// SpamHandler represents a concrete handler
type SpamHandler struct {
	BaseHandler
}

// Handle implements the Handler interface
func (h *SpamHandler) Handle(request string) string {
	if request == "spam" {
		return "SpamHandler: Handled spam"
	}
	return h.BaseHandler.Handle(request)
}

// FanHandler represents a concrete handler
type FanHandler struct {
	BaseHandler
}

// Handle implements the Handler interface
func (h *FanHandler) Handle(request string) string {
	if request == "fan" {
		return "FanHandler: Handled fan mail"
	}
	return h.BaseHandler.Handle(request)
}

// ComplaintHandler represents a concrete handler
type ComplaintHandler struct {
	BaseHandler
}

// Handle implements the Handler interface
func (h *ComplaintHandler) Handle(request string) string {
	if request == "complaint" {
		return "ComplaintHandler: Handled complaint"
	}
	return h.BaseHandler.Handle(request)
}

// NewHandler creates a new handler chain
func NewHandler() Handler {
	spam := &SpamHandler{}
	fan := &FanHandler{}
	complaint := &ComplaintHandler{}

	spam.SetNext(fan)
	fan.SetNext(complaint)

	return spam
}

// ChainOfResponsibilityDemo demonstrates the Chain of Responsibility pattern
func ChainOfResponsibilityDemo() {
	handler := NewHandler()

	fmt.Println("--- Chain of Responsibility Pattern Demo ---")
	
	requests := []string{"spam", "fan", "complaint", "other"}
	
	for _, request := range requests {
		result := handler.Handle(request)
		if result == "" {
			fmt.Printf("No handler found for request: %s\n", request)
		} else {
			fmt.Println(result)
		}
	}
} 