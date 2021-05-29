package main

import (
	"fmt"
	"os"
)

func main() {
	scooter, scooterErr := getTransport("scooter")
	if scooterErr != nil {
		fmt.Printf("Scooter creation failed: %s \n", scooterErr.Error())
		os.Exit(1)
	}
	fmt.Printf("Scooter: %+v \n", scooter)

	quadcopter, quadErr := getTransport("quadcopter")
	if quadErr != nil {
		fmt.Printf("Quadcopter creation failed: %s \n", quadErr.Error())
		os.Exit(1)
	}
	fmt.Printf("Quadcopter: %+v \n", quadcopter)
}
