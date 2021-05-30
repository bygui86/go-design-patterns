package display

import "fmt"

// CurrentConditionDisplay implements commons.Observer and commons.DisplayElement interfaces
type CurrentConditionDisplay struct {
	temperature float32
	humidity    float32
	pressure    float32
}

func NewCurrentConditionDisplay() *CurrentConditionDisplay {
	return &CurrentConditionDisplay{}
}

/*
	This implements Observer and DisplayElement interfaces to get changes from WeatherData
	and show the information based on its functionality respectively
*/
func (d *CurrentConditionDisplay) Update(temp float32, humidity float32, pressure float32) {
	d.temperature = temp
	d.humidity = humidity
	d.pressure = pressure
	// When update() is called, we save the measurements and call display() to show the required information
	d.Display()
}

func (d *CurrentConditionDisplay) Display() {
	fmt.Printf("Current conditions: Temperature:%.2f, Humidity:%.2f and Pressure:%.2f\n", d.temperature, d.humidity, d.pressure)
}
