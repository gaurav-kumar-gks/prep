package structural

import "fmt"

// Device defines the implementation interface
type Device interface {
	TurnOn()
	TurnOff()
	SetVolume(volume int)
	SetChannel(channel int)
}

// TV is a concrete implementation of Device
type TV struct {
	isOn    bool
	volume  int
	channel int
}

// NewTV creates a new TV
func NewTV() *TV {
	return &TV{
		isOn:    false,
		volume:  30,
		channel: 1,
	}
}

// TurnOn implements the Device interface
func (t *TV) TurnOn() {
	t.isOn = true
	fmt.Println("TV is turned on")
}

// TurnOff implements the Device interface
func (t *TV) TurnOff() {
	t.isOn = false
	fmt.Println("TV is turned off")
}

// SetVolume implements the Device interface
func (t *TV) SetVolume(volume int) {
	t.volume = volume
	fmt.Println("TV volume set to", volume)
}

// SetChannel implements the Device interface
func (t *TV) SetChannel(channel int) {
	t.channel = channel
	fmt.Println("TV channel set to", channel)
}

// Radio is a concrete implementation of Device
type Radio struct {
	isOn    bool
	volume  int
	channel int
}

// NewRadio creates a new Radio
func NewRadio() *Radio {
	return &Radio{
		isOn:    false,
		volume:  30,
		channel: 1,
	}
}

// TurnOn implements the Device interface
func (r *Radio) TurnOn() {
	r.isOn = true
	fmt.Println("Radio is turned on")
}

// TurnOff implements the Device interface
func (r *Radio) TurnOff() {
	r.isOn = false
	fmt.Println("Radio is turned off")
}

// SetVolume implements the Device interface
func (r *Radio) SetVolume(volume int) {
	r.volume = volume
	fmt.Println("Radio volume set to", volume)
}

// SetChannel implements the Device interface
func (r *Radio) SetChannel(channel int) {
	r.channel = channel
	fmt.Println("Radio channel set to", channel)
}

// RemoteControl is the abstraction
type RemoteControl struct {
	device Device
}

// NewRemoteControl creates a new RemoteControl
func NewRemoteControl(device Device) *RemoteControl {
	return &RemoteControl{device: device}
}

// TurnOn turns on the device
func (r *RemoteControl) TurnOn() {
	r.device.TurnOn()
}

// TurnOff turns off the device
func (r *RemoteControl) TurnOff() {
	r.device.TurnOff()
}

// VolumeUp increases the volume
func (r *RemoteControl) VolumeUp() {
	// Implementation specific to the device type
	switch r.device.(type) {
	case *TV:
		r.device.SetVolume(50)
	default:
		r.device.SetVolume(40)
	}
}

// VolumeDown decreases the volume
func (r *RemoteControl) VolumeDown() {
	// Implementation specific to the device type
	switch r.device.(type) {
	case *TV:
		r.device.SetVolume(10)
	default:
		r.device.SetVolume(20)
	}
}

// ChannelUp increases the channel
func (r *RemoteControl) ChannelUp() {
	// Implementation specific to the device type
	switch r.device.(type) {
	case *TV:
		r.device.SetChannel(10)
	default:
		r.device.SetChannel(5)
	}
}

// ChannelDown decreases the channel
func (r *RemoteControl) ChannelDown() {
	r.device.SetChannel(1)
}

// BridgeDemo demonstrates the Bridge pattern
func BridgeDemo() {
	fmt.Println("Testing TV with basic remote:")
	tv := NewTV()
	tvRemote := NewRemoteControl(tv)
	tvRemote.TurnOn()
	tvRemote.VolumeUp()
	tvRemote.ChannelUp()
	tvRemote.TurnOff()
	
	fmt.Println("\nTesting Radio with basic remote:")
	radio := NewRadio()
	radioRemote := NewRemoteControl(radio)
	radioRemote.TurnOn()
	radioRemote.VolumeDown()
	radioRemote.ChannelDown()
	radioRemote.TurnOff()
} 