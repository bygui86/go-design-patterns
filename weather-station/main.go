package main

import (
	"github.com/bygui86/go-design-patterns/weather-station/data"
	"github.com/bygui86/go-design-patterns/weather-station/display"
)

func main() {
	// Creating the weather object
	weatherData := data.NewWeatherData()

	// Creating the displays and registering as observer to the Subject
	currentConditionDisplay := display.NewCurrentConditionDisplay()
	weatherData.RegisterObserver(currentConditionDisplay)

	statisticsDisplay := display.NewStatisticsDisplay()
	weatherData.RegisterObserver(statisticsDisplay)

	// Simulate new weather measurements
	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(77, 90, 29.2)

	// Deregistering the statisticsDisplay. Change in measurements won't notify this observer.
	weatherData.DeregisterObserver(statisticsDisplay)

	weatherData.SetMeasurements(84, 80, 32.6)
}
