package data

import "github.com/bygui86/go-design-patterns/weather-station/commons"

// WeatherData implements subject interface
type WeatherData struct {
	observers   map[commons.Observer]bool
	temperature float32
	humidity    float32
	pressure    float32
}

func NewWeatherData() *WeatherData {
	return &WeatherData{
		observers: make(map[commons.Observer]bool),
	}
}

// WeatherData now implements the Subject interface
func (w *WeatherData) RegisterObserver(obs commons.Observer) {
	// When an observer registers, we add it in the map with value set as true
	w.observers[obs] = true
}

func (w *WeatherData) DeregisterObserver(obs commons.Observer) {
	// When an observer deregisters itself, we remove it from the map after checking if the value exists
	if _, ok := w.observers[obs]; ok {
		delete(w.observers, obs)
	}
}

func (w *WeatherData) NotifyObservers() {
	// This is where we tell all the observers about the state by calling the common update method
	for obs := range w.observers {
		obs.Update(w.temperature, w.humidity, w.pressure)
	}
}

// Dummy method to test our display elements. Setting the measurements not via a device.
func (w *WeatherData) SetMeasurements(temp float32, humidity float32, pressure float32) {
	w.temperature = temp
	w.humidity = humidity
	w.pressure = pressure

	w.MeasurementsChanged()
}

// We notify the Observers when we get updated measurements from the Weather Station
func (w *WeatherData) MeasurementsChanged() {
	w.NotifyObservers()
}
