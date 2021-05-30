package commons

type Subject interface {
	// Both of these methods take an Observer as argument
	RegisterObserver(Observer)
	DeregisterObserver(Observer)

	// This method is called to notify all the observers when the Subject's state has changed
	NotifyObservers()
}

/*
	Observer interface is implemented by all observers, here we're passing the measurements to the observers
	from the Subject when a weather measurement changes
*/
type Observer interface {
	Update(temp float32, humidity float32, pressure float32)
}
