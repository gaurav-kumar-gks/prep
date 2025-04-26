package behavioral

import "fmt"

// Mediator defines the mediator interface
type Mediator interface {
	Notify(sender Component, event string)
}

// Component defines the component interface
type Component interface {
	SetMediator(mediator Mediator)
}

// BaseComponent provides default implementation
type BaseComponent struct {
	mediator Mediator
}

// SetMediator sets the mediator
func (c *BaseComponent) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

// Button represents a concrete component
type Button struct {
	BaseComponent
}

// Click simulates a button click
func (b *Button) Click() {
	fmt.Println("Button clicked")
	b.mediator.Notify(b, "click")
}

// TextBox represents a concrete component
type TextBox struct {
	BaseComponent
	text string
}

// SetText sets the text
func (t *TextBox) SetText(text string) {
	t.text = text
	fmt.Printf("TextBox text set to: %s\n", text)
	t.mediator.Notify(t, "textChanged")
}

// GetText gets the text
func (t *TextBox) GetText() string {
	return t.text
}

// CheckBox represents a concrete component
type CheckBox struct {
	BaseComponent
	checked bool
}

// Check checks the checkbox
func (c *CheckBox) Check() {
	c.checked = true
	fmt.Println("CheckBox checked")
	c.mediator.Notify(c, "checked")
}

// Uncheck unchecks the checkbox
func (c *CheckBox) Uncheck() {
	c.checked = false
	fmt.Println("CheckBox unchecked")
	c.mediator.Notify(c, "unchecked")
}

// IsChecked returns whether the checkbox is checked
func (c *CheckBox) IsChecked() bool {
	return c.checked
}

// Dialog represents a concrete mediator
type Dialog struct {
	button   *Button
	textBox  *TextBox
	checkBox *CheckBox
}

// NewDialog creates a new Dialog
func NewDialog() *Dialog {
	dialog := &Dialog{}
	
	dialog.button = &Button{}
	dialog.textBox = &TextBox{}
	dialog.checkBox = &CheckBox{}
	
	dialog.button.SetMediator(dialog)
	dialog.textBox.SetMediator(dialog)
	dialog.checkBox.SetMediator(dialog)
	
	return dialog
}

// Notify implements the Mediator interface
func (d *Dialog) Notify(sender Component, event string) {
	switch event {
	case "click":
		if d.checkBox.IsChecked() {
			d.textBox.SetText("Button clicked and checkbox is checked")
		} else {
			d.textBox.SetText("Button clicked and checkbox is unchecked")
		}
	case "checked":
		d.textBox.SetText("Checkbox is now checked")
	case "unchecked":
		d.textBox.SetText("Checkbox is now unchecked")
	case "textChanged":
		if d.textBox.GetText() == "" {
			d.button.Click()
		}
	}
}

// MediatorDemo demonstrates the Mediator pattern
func MediatorDemo() {
	dialog := NewDialog()

	fmt.Println("--- Mediator Pattern Demo ---")
	
	fmt.Println("\n1. Click button with unchecked checkbox:")
	dialog.button.Click()
	
	fmt.Println("\n2. Check the checkbox:")
	dialog.checkBox.Check()
	
	fmt.Println("\n3. Click button with checked checkbox:")
	dialog.button.Click()
	
	fmt.Println("\n4. Uncheck the checkbox:")
	dialog.checkBox.Uncheck()
	
	fmt.Println("\n5. Set text box text:")
	dialog.textBox.SetText("Hello, Mediator!")
} 