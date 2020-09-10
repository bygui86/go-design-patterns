package main

import (
	"fmt"

	"github.com/bygui86/go-design-patterns/bus-service/internal"
)

func main() {
	fmt.Println("Starting simulation")
	expressLine := internal.NewBus("Express Line", "")

	s1 := internal.BusStop{Name: "Downtown"}
	s2 := internal.BusStop{Name: "The University"}
	s3 := internal.BusStop{Name: "The Village"}

	expressLine.AddStop(&s1)
	expressLine.AddStop(&s2)
	expressLine.AddStop(&s3)

	john := internal.Prospect{
		SSN:         "12345612-22",
		Destination: &s2,
	}
	betty := internal.Prospect{
		SSN:         "11223322-67",
		Destination: &s3,
	}
	s1.NotifyProspectArrival(john)
	s1.NotifyProspectArrival(betty)

	for expressLine.Go() {
		expressLine.VisitPassengers(func(p internal.Passenger) {
			fmt.Printf("    Passenger with SSN %q is heading to %q\n", p.SSN, p.Destination.Name)
		})
	}
	fmt.Println("Simulation done")
}
