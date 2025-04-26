package behavioral

import "fmt"

// Command defines the command interface
type Command interface {
	Execute()
	Undo()
}

// Light represents a receiver
type Light struct {
	location string
}

// NewLight creates a new Light
func NewLight(location string) *Light {
	return &Light{location: location}
}

// TurnOn turns on the light
func (l *Light) TurnOn() {
	fmt.Printf("Light is on in %s\n", l.location)
}

// TurnOff turns off the light
func (l *Light) TurnOff() {
	fmt.Printf("Light is off in %s\n", l.location)
}

// LightOnCommand is a concrete command
type LightOnCommand struct {
	light *Light
}

// NewLightOnCommand creates a new LightOnCommand
func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

// Execute implements the Command interface
func (c *LightOnCommand) Execute() {
	c.light.TurnOn()
}

// Undo implements the Command interface
func (c *LightOnCommand) Undo() {
	c.light.TurnOff()
}

// LightOffCommand is a concrete command
type LightOffCommand struct {
	light *Light
}

// NewLightOffCommand creates a new LightOffCommand
func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

// Execute implements the Command interface
func (c *LightOffCommand) Execute() {
	c.light.TurnOff()
}

// Undo implements the Command interface
func (c *LightOffCommand) Undo() {
	c.light.TurnOn()
}

// RemoteControl represents the invoker
type RemoteControl struct {
	onCommands  []Command
	offCommands []Command
	undoCommand Command
}

// NewRemoteControl creates a new RemoteControl
func NewRemoteControl() *RemoteControl {
	return &RemoteControl{
		onCommands:  make([]Command, 7),
		offCommands: make([]Command, 7),
	}
}

// SetCommand sets a command for a slot
func (r *RemoteControl) SetCommand(slot int, onCommand Command, offCommand Command) {
	r.onCommands[slot] = onCommand
	r.offCommands[slot] = offCommand
}

// OnButtonWasPushed handles the on button press
func (r *RemoteControl) OnButtonWasPushed(slot int) {
	if r.onCommands[slot] != nil {
		r.onCommands[slot].Execute()
		r.undoCommand = r.onCommands[slot]
	}
}

// OffButtonWasPushed handles the off button press
func (r *RemoteControl) OffButtonWasPushed(slot int) {
	if r.offCommands[slot] != nil {
		r.offCommands[slot].Execute()
		r.undoCommand = r.offCommands[slot]
	}
}

// UndoButtonWasPushed handles the undo button press
func (r *RemoteControl) UndoButtonWasPushed() {
	if r.undoCommand != nil {
		r.undoCommand.Undo()
	}
}

// CommandDemo demonstrates the Command pattern
func CommandDemo() {
	// Create receivers
	livingRoomLight := NewLight("Living Room")
	kitchenLight := NewLight("Kitchen")

	// Create commands
	livingRoomLightOn := NewLightOnCommand(livingRoomLight)
	livingRoomLightOff := NewLightOffCommand(livingRoomLight)
	kitchenLightOn := NewLightOnCommand(kitchenLight)
	kitchenLightOff := NewLightOffCommand(kitchenLight)

	// Create invoker
	remote := NewRemoteControl()

	// Set commands
	remote.SetCommand(0, livingRoomLightOn, livingRoomLightOff)
	remote.SetCommand(1, kitchenLightOn, kitchenLightOff)

	// Use remote
	fmt.Println("--- Remote Control Demo ---")
	remote.OnButtonWasPushed(0)
	remote.OffButtonWasPushed(0)
	remote.UndoButtonWasPushed()

	remote.OnButtonWasPushed(1)
	remote.OffButtonWasPushed(1)
	remote.UndoButtonWasPushed()
} 