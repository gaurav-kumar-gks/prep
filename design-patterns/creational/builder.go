package creational

import "fmt"

// House represents the product
type House struct {
	windowType string
	doorType   string
	floor      int
}

// HouseBuilder defines the builder interface
type HouseBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() *House
}

// IglooBuilder is a concrete builder for Igloo houses
type IglooBuilder struct {
	house *House
}

// NewIglooBuilder creates a new IglooBuilder
func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{house: &House{}}
}

// setWindowType sets the window type for Igloo
func (b *IglooBuilder) setWindowType() {
	b.house.windowType = "Snow Window"
}

// setDoorType sets the door type for Igloo
func (b *IglooBuilder) setDoorType() {
	b.house.doorType = "Snow Door"
}

// setNumFloor sets the number of floors for Igloo
func (b *IglooBuilder) setNumFloor() {
	b.house.floor = 1
}

// getHouse returns the built house
func (b *IglooBuilder) getHouse() *House {
	return b.house
}

// TipiBuilder is a concrete builder for Tipi houses
type TipiBuilder struct {
	house *House
}

// NewTipiBuilder creates a new TipiBuilder
func NewTipiBuilder() *TipiBuilder {
	return &TipiBuilder{house: &House{}}
}

// setWindowType sets the window type for Tipi
func (b *TipiBuilder) setWindowType() {
	b.house.windowType = "Wooden Window"
}

// setDoorType sets the door type for Tipi
func (b *TipiBuilder) setDoorType() {
	b.house.doorType = "Wooden Door"
}

// setNumFloor sets the number of floors for Tipi
func (b *TipiBuilder) setNumFloor() {
	b.house.floor = 1
}

// getHouse returns the built house
func (b *TipiBuilder) getHouse() *House {
	return b.house
}

// NormalBuilder is a concrete builder for Normal houses
type NormalBuilder struct {
	house *House
}

// NewNormalBuilder creates a new NormalBuilder
func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{house: &House{}}
}

// setWindowType sets the window type for Normal house
func (b *NormalBuilder) setWindowType() {
	b.house.windowType = "Glass Window"
}

// setDoorType sets the door type for Normal house
func (b *NormalBuilder) setDoorType() {
	b.house.doorType = "Wooden Door"
}

// setNumFloor sets the number of floors for Normal house
func (b *NormalBuilder) setNumFloor() {
	b.house.floor = 2
}

// getHouse returns the built house
func (b *NormalBuilder) getHouse() *House {
	return b.house
}

// Director orchestrates the building process
type Director struct {
	builder HouseBuilder
}

// NewDirector creates a new Director
func NewDirector(builder HouseBuilder) *Director {
	return &Director{builder: builder}
}

// setBuilder sets the builder
func (d *Director) setBuilder(builder HouseBuilder) {
	d.builder = builder
}

// constructHouse constructs a house
func (d *Director) constructHouse() *House {
	d.builder.setWindowType()
	d.builder.setDoorType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

// BuilderDemo demonstrates the Builder pattern
func BuilderDemo() {
	normalBuilder := NewNormalBuilder()
	iglooBuilder := NewIglooBuilder()
	tipiBuilder := NewTipiBuilder()

	director := NewDirector(normalBuilder)
	normalHouse := director.constructHouse()
	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.constructHouse()
	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Floor: %d\n", iglooHouse.floor)

	director.setBuilder(tipiBuilder)
	tipiHouse := director.constructHouse()
	fmt.Printf("\nTipi House Door Type: %s\n", tipiHouse.doorType)
	fmt.Printf("Tipi House Window Type: %s\n", tipiHouse.windowType)
	fmt.Printf("Tipi House Floor: %d\n", tipiHouse.floor)
} 