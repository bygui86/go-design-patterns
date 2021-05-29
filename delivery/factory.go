package main

import "fmt"

// Defines an interface for transport objects
type iTransport interface {
	// SetName sets the name of a transport
	SetName(n string)
	// GetName returns the name of transport
	GetName() string
	// SetSpeed sets the speed of a transport
	SetSpeed(s uint)
	// GetSpeed returns the speed of transport
	GetSpeed() uint
}

func getTransport(transportType string) (iTransport, error) {
	if transportType == "scooter" {
		return newScooter(), nil
	}

	if transportType == "quadcopter" {
		return newQuadcopter(), nil
	}

	return nil, fmt.Errorf("unknown type %s", transportType)
}
