package structural

import "fmt"

// CPU represents the CPU subsystem
type CPU struct{}

// Freeze freezes the CPU
func (c *CPU) Freeze() {
	fmt.Println("CPU: Freezing...")
}

// Jump jumps to a position
func (c *CPU) Jump(position int64) {
	fmt.Println("CPU: Jumping to position", position)
}

// Execute executes the CPU
func (c *CPU) Execute() {
	fmt.Println("CPU: Executing...")
}

// Memory represents the Memory subsystem
type Memory struct{}

// Load loads data into memory
func (m *Memory) Load(position int64, data []byte) {
	fmt.Println("Memory: Loading data at position", position)
}

// HardDrive represents the HardDrive subsystem
type HardDrive struct{}

// Read reads data from the hard drive
func (h *HardDrive) Read(lba int64, size int) []byte {
	fmt.Println("HardDrive: Reading", size, "bytes from LBA", lba)
	return make([]byte, size)
}

// ComputerFacade is the facade for the computer subsystem
type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

// NewComputerFacade creates a new ComputerFacade
func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

// Start starts the computer
func (c *ComputerFacade) Start() {
	fmt.Println("Computer: Starting...")
	c.cpu.Freeze()
	c.memory.Load(0, c.hardDrive.Read(0, 1024))
	c.cpu.Jump(0)
	c.cpu.Execute()
	fmt.Println("Computer: Started successfully")
}

// Shutdown shuts down the computer
func (c *ComputerFacade) Shutdown() {
	fmt.Println("Computer: Shutting down...")
	// Add shutdown logic here
	fmt.Println("Computer: Shutdown complete")
}

// FacadeDemo demonstrates the Facade pattern
func FacadeDemo() {
	computer := NewComputerFacade()
	
	fmt.Println("Starting computer:")
	computer.Start()
	
	fmt.Println("\nShutting down computer:")
	computer.Shutdown()
} 