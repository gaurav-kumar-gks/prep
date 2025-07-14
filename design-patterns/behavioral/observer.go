package behavioral

import (
	"fmt"
	"sync"
)

// Subject defines the subject interface
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

// Observer defines the observer interface
type Observer interface {
	Update(temperature float64, humidity float64, pressure float64)
}

// WeatherData represents the concrete subject
type WeatherData struct {
	observers   []Observer
	temperature float64
	humidity    float64
	pressure    float64
	mutex       sync.Mutex
}

// NewWeatherData creates a new WeatherData
func NewWeatherData() *WeatherData {
	return &WeatherData{
		observers: make([]Observer, 0),
	}
}

// RegisterObserver implements the Subject interface
func (w *WeatherData) RegisterObserver(observer Observer) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.observers = append(w.observers, observer)
}

// RemoveObserver implements the Subject interface
func (w *WeatherData) RemoveObserver(observer Observer) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	for i, obs := range w.observers {
		if obs == observer {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers implements the Subject interface
func (w *WeatherData) NotifyObservers() {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	for _, observer := range w.observers {
		observer.Update(w.temperature, w.humidity, w.pressure)
	}
}

// SetMeasurements updates the weather measurements
func (w *WeatherData) SetMeasurements(temperature float64, humidity float64, pressure float64) {
	w.mutex.Lock()
	w.temperature = temperature
	w.humidity = humidity
	w.pressure = pressure
	w.mutex.Unlock()
	w.NotifyObservers()
}

// CurrentConditionsDisplay represents a concrete observer
type CurrentConditionsDisplay struct {
	temperature float64
	humidity    float64
	weatherData Subject
}

// NewCurrentConditionsDisplay creates a new CurrentConditionsDisplay
func NewCurrentConditionsDisplay(weatherData Subject) *CurrentConditionsDisplay {
	display := &CurrentConditionsDisplay{weatherData: weatherData}
	weatherData.RegisterObserver(display)
	return display
}

// Update implements the Observer interface
func (d *CurrentConditionsDisplay) Update(temperature float64, humidity float64, pressure float64) {
	d.temperature = temperature
	d.humidity = humidity
	d.Display()
}

// Display shows the current conditions
func (d *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current conditions: %.1f°F degrees and %.1f%% humidity\n",
		d.temperature, d.humidity)
}

// StatisticsDisplay represents a concrete observer
type StatisticsDisplay struct {
	temperatures []float64
	weatherData  Subject
}

// NewStatisticsDisplay creates a new StatisticsDisplay
func NewStatisticsDisplay(weatherData Subject) *StatisticsDisplay {
	display := &StatisticsDisplay{
		temperatures: make([]float64, 0),
		weatherData:  weatherData,
	}
	weatherData.RegisterObserver(display)
	return display
}

// Update implements the Observer interface
func (d *StatisticsDisplay) Update(temperature float64, humidity float64, pressure float64) {
	d.temperatures = append(d.temperatures, temperature)
	d.Display()
}

// Display shows the statistics
func (d *StatisticsDisplay) Display() {
	if len(d.temperatures) == 0 {
		return
	}

	var sum float64
	for _, temp := range d.temperatures {
		sum += temp
	}
	avg := sum / float64(len(d.temperatures))

	fmt.Printf("Avg/Max/Min temperature = %.1f/%.1f/%.1f\n",
		avg, d.temperatures[len(d.temperatures)-1], d.temperatures[0])
}

// ObserverDemo demonstrates the Observer pattern
func ObserverDemo() {
	weatherData := NewWeatherData()

	// currentDisplay := NewCurrentConditionsDisplay(weatherData)
	// statisticsDisplay := NewStatisticsDisplay(weatherData)

	fmt.Println("--- Weather Station Demo ---")
	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.2)
} 