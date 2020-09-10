package internal

// BusStop represents a place where a Bus can stop and signal to prospects (future passengers)
// that they may board.
type BusStop struct {
	Name      string
	prospects []Prospect
	busses    []Bus
}

// Prospect is a potential Passenger. Prospects wait at BusStops to board Buses.
type Prospect struct {
	SSN         string
	Destination *BusStop
}

// BUS STOP

// NotifyProspectArrival is called whenever a prospect arrives at Busstop.
func (b *BusStop) NotifyProspectArrival(p Prospect) {
	b.prospects = append(b.prospects, p)

	// Notify all Busses on this route.
	for _, bus := range b.busses {
		if bus.StopsAt(p.Destination) {
			bus.NotifyBoardingIntent(b)
		}
	}
}

// NotifyBusArrival is called by Bus upon arrival.
func (b *BusStop) NotifyBusArrival(bus *Bus) {
	for _, p := range b.prospects {
		if bus.StopsAt(p.Destination) {
			pas := p.ToPassenger()
			bus.Board(&pas, bus.company.GetPricing())
		}
	}
}

// Equals returns true if the given BusStop is the same as the receiver.
func (b *BusStop) Equals(busStop *BusStop) bool {
	return b.Name == busStop.Name
}

// PROSPECT

// ToPassenger returns a Passenger with the same SSN as his or her Prospect.
func (p Prospect) ToPassenger() Passenger {
	return Passenger{
		SSN:         p.SSN,
		Destination: p.Destination,
	}
}
