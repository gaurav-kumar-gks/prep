package behavioral

import (
	"fmt"
	"sync"
)

// State defines the state interface
type State interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	Dispense()
}

// GumballMachine represents the context
type GumballMachine struct {
	soldOutState    State
	noQuarterState  State
	hasQuarterState State
	soldState       State
	winnerState     State

	state     State
	count     int
	location  string
	mutex     sync.Mutex
}

// NewGumballMachine creates a new GumballMachine
func NewGumballMachine(location string, count int) *GumballMachine {
	machine := &GumballMachine{
		location: location,
		count:    count,
	}

	machine.soldOutState = &SoldOutState{machine: machine}
	machine.noQuarterState = &NoQuarterState{machine: machine}
	machine.hasQuarterState = &HasQuarterState{machine: machine}
	machine.soldState = &SoldState{machine: machine}
	machine.winnerState = &WinnerState{machine: machine}

	if count > 0 {
		machine.state = machine.noQuarterState
	} else {
		machine.state = machine.soldOutState
	}

	return machine
}

// InsertQuarter inserts a quarter
func (m *GumballMachine) InsertQuarter() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.state.InsertQuarter()
}

// EjectQuarter ejects a quarter
func (m *GumballMachine) EjectQuarter() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.state.EjectQuarter()
}

// TurnCrank turns the crank
func (m *GumballMachine) TurnCrank() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.state.TurnCrank()
	m.state.Dispense()
}

// SetState sets the state
func (m *GumballMachine) SetState(state State) {
	m.state = state
}

// ReleaseBall releases a gumball
func (m *GumballMachine) ReleaseBall() {
	if m.count > 0 {
		m.count--
		fmt.Println("A gumball comes rolling out the slot...")
	}
}

// GetCount returns the number of gumballs
func (m *GumballMachine) GetCount() int {
	return m.count
}

// Refill refills the machine
func (m *GumballMachine) Refill(count int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.count = count
	m.state = m.noQuarterState
}

// GetState returns the current state
func (m *GumballMachine) GetState() State {
	return m.state
}

// GetLocation returns the location
func (m *GumballMachine) GetLocation() string {
	return m.location
}

// SoldOutState represents the sold out state
type SoldOutState struct {
	machine *GumballMachine
}

// InsertQuarter implements the State interface
func (s *SoldOutState) InsertQuarter() {
	fmt.Println("You can't insert a quarter, the machine is sold out")
}

// EjectQuarter implements the State interface
func (s *SoldOutState) EjectQuarter() {
	fmt.Println("You can't eject, you haven't inserted a quarter yet")
}

// TurnCrank implements the State interface
func (s *SoldOutState) TurnCrank() {
	fmt.Println("You turned, but there are no gumballs")
}

// Dispense implements the State interface
func (s *SoldOutState) Dispense() {
	fmt.Println("No gumball dispensed")
}

// NoQuarterState represents the no quarter state
type NoQuarterState struct {
	machine *GumballMachine
}

// InsertQuarter implements the State interface
func (s *NoQuarterState) InsertQuarter() {
	fmt.Println("You inserted a quarter")
	s.machine.SetState(s.machine.hasQuarterState)
}

// EjectQuarter implements the State interface
func (s *NoQuarterState) EjectQuarter() {
	fmt.Println("You haven't inserted a quarter")
}

// TurnCrank implements the State interface
func (s *NoQuarterState) TurnCrank() {
	fmt.Println("You turned, but there's no quarter")
}

// Dispense implements the State interface
func (s *NoQuarterState) Dispense() {
	fmt.Println("You need to pay first")
}

// HasQuarterState represents the has quarter state
type HasQuarterState struct {
	machine *GumballMachine
}

// InsertQuarter implements the State interface
func (s *HasQuarterState) InsertQuarter() {
	fmt.Println("You can't insert another quarter")
}

// EjectQuarter implements the State interface
func (s *HasQuarterState) EjectQuarter() {
	fmt.Println("Quarter returned")
	s.machine.SetState(s.machine.noQuarterState)
}

// TurnCrank implements the State interface
func (s *HasQuarterState) TurnCrank() {
	fmt.Println("You turned...")
	if s.machine.GetCount() > 1 && s.machine.GetCount()%10 == 0 {
		s.machine.SetState(s.machine.winnerState)
	} else {
		s.machine.SetState(s.machine.soldState)
	}
}

// Dispense implements the State interface
func (s *HasQuarterState) Dispense() {
	fmt.Println("No gumball dispensed")
}

// SoldState represents the sold state
type SoldState struct {
	machine *GumballMachine
}

// InsertQuarter implements the State interface
func (s *SoldState) InsertQuarter() {
	fmt.Println("Please wait, we're already giving you a gumball")
}

// EjectQuarter implements the State interface
func (s *SoldState) EjectQuarter() {
	fmt.Println("Sorry, you already turned the crank")
}

// TurnCrank implements the State interface
func (s *SoldState) TurnCrank() {
	fmt.Println("Turning twice doesn't get you another gumball!")
}

// Dispense implements the State interface
func (s *SoldState) Dispense() {
	s.machine.ReleaseBall()
	if s.machine.GetCount() > 0 {
		s.machine.SetState(s.machine.noQuarterState)
	} else {
		fmt.Println("Oops, out of gumballs!")
		s.machine.SetState(s.machine.soldOutState)
	}
}

// WinnerState represents the winner state
type WinnerState struct {
	machine *GumballMachine
}

// InsertQuarter implements the State interface
func (s *WinnerState) InsertQuarter() {
	fmt.Println("Please wait, we're already giving you a gumball")
}

// EjectQuarter implements the State interface
func (s *WinnerState) EjectQuarter() {
	fmt.Println("Sorry, you already turned the crank")
}

// TurnCrank implements the State interface
func (s *WinnerState) TurnCrank() {
	fmt.Println("Turning twice doesn't get you another gumball!")
}

// Dispense implements the State interface
func (s *WinnerState) Dispense() {
	fmt.Println("YOU'RE A WINNER! You get two gumballs for your quarter")
	s.machine.ReleaseBall()
	if s.machine.GetCount() == 0 {
		s.machine.SetState(s.machine.soldOutState)
	} else {
		s.machine.ReleaseBall()
		if s.machine.GetCount() > 0 {
			s.machine.SetState(s.machine.noQuarterState)
		} else {
			fmt.Println("Oops, out of gumballs!")
			s.machine.SetState(s.machine.soldOutState)
		}
	}
}

// StateDemo demonstrates the State pattern
func StateDemo() {
	gumballMachine := NewGumballMachine("Mighty Gumball, Inc.", 5)

	fmt.Println("--- Gumball Machine Demo ---")
	fmt.Printf("Machine location: %s\n", gumballMachine.GetLocation())
	fmt.Printf("Current inventory: %d gumballs\n", gumballMachine.GetCount())

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	fmt.Printf("\nCurrent inventory: %d gumballs\n", gumballMachine.GetCount())

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	fmt.Printf("\nCurrent inventory: %d gumballs\n", gumballMachine.GetCount())
} 