package creational

import "fmt"

// Button defines the interface for buttons
type Button interface {
	Render()
	OnClick()
}

// Checkbox defines the interface for checkboxes
type Checkbox interface {
	Render()
	OnCheck()
}

// WindowsButton is a concrete button for Windows
type WindowsButton struct{}

// Render renders the Windows button
func (b *WindowsButton) Render() {
	fmt.Println("Rendering Windows button")
}

// OnClick handles the click event for Windows button
func (b *WindowsButton) OnClick() {
	fmt.Println("Windows button clicked")
}

// WindowsCheckbox is a concrete checkbox for Windows
type WindowsCheckbox struct{}

// Render renders the Windows checkbox
func (c *WindowsCheckbox) Render() {
	fmt.Println("Rendering Windows checkbox")
}

// OnCheck handles the check event for Windows checkbox
func (c *WindowsCheckbox) OnCheck() {
	fmt.Println("Windows checkbox checked")
}

// MacButton is a concrete button for Mac
type MacButton struct{}

// Render renders the Mac button
func (b *MacButton) Render() {
	fmt.Println("Rendering Mac button")
}

// OnClick handles the click event for Mac button
func (b *MacButton) OnClick() {
	fmt.Println("Mac button clicked")
}

// MacCheckbox is a concrete checkbox for Mac
type MacCheckbox struct{}

// Render renders the Mac checkbox
func (c *MacCheckbox) Render() {
	fmt.Println("Rendering Mac checkbox")
}

// OnCheck handles the check event for Mac checkbox
func (c *MacCheckbox) OnCheck() {
	fmt.Println("Mac checkbox checked")
}

// GUIFactory defines the abstract factory interface
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// WindowsFactory is a concrete factory for Windows
type WindowsFactory struct{}

// CreateButton creates a Windows button
func (f *WindowsFactory) CreateButton() Button {
	return &WindowsButton{}
}

// CreateCheckbox creates a Windows checkbox
func (f *WindowsFactory) CreateCheckbox() Checkbox {
	return &WindowsCheckbox{}
}

// MacFactory is a concrete factory for Mac
type MacFactory struct{}

// CreateButton creates a Mac button
func (f *MacFactory) CreateButton() Button {
	return &MacButton{}
}

// CreateCheckbox creates a Mac checkbox
func (f *MacFactory) CreateCheckbox() Checkbox {
	return &MacCheckbox{}
}

// Application represents the client
type Application struct {
	factory GUIFactory
}

// NewApplication creates a new application
func NewApplication(factory GUIFactory) *Application {
	return &Application{factory: factory}
}

// CreateUI creates the UI
func (a *Application) CreateUI() {
	button := a.factory.CreateButton()
	checkbox := a.factory.CreateCheckbox()

	button.Render()
	checkbox.Render()
}

// AbstractFactoryDemo demonstrates the Abstract Factory pattern
func AbstractFactoryDemo() {
	fmt.Println("Creating Windows UI:")
	app := NewApplication(&WindowsFactory{})
	app.CreateUI()

	fmt.Println("\nCreating Mac UI:")
	app = NewApplication(&MacFactory{})
	app.CreateUI()
} 