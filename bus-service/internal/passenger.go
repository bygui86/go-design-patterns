package internal

import (
	"fmt"
	"strconv"
)

// SeniorAge is the minimum age from which a Passenger is considered a senior to the BusCompany.
const SeniorAge = 65

// Passenger represents a bus passenger, uniquely identified by their SSN.
type Passenger struct {
	SSN            string
	SeatNumber     uint8
	Destination    *BusStop
	HasValidTicket bool
}

// Passengers represents a set of Passengers, using their SSN as key.
// Notice how the responsibility of handling passengers has shifted from Bus to Passengers
type Passengers map[string]Passenger

// NewPassengerSet returns an empty set of Passengers, ready to use.
func NewPassengerSet() Passengers {
	return make(map[string]Passenger)
}

// Add adds a Passenger to Passengers. The Passenger will be overwritten if exists.
func (p Passengers) Add(newP Passenger) {
	p[newP.SSN] = newP
}

// Remove removes a Passenger from Passengers.
func (p Passengers) Remove(delP Passenger) {
	delete(p, delP.SSN)
}

// Visit calls visitor once for every Passenger in the set.
func (p Passengers) Visit(visitor func(Passenger)) {
	for _, one := range p {
		visitor(one)
	}
}

// Find returns the Passenger with the given SSN. If none was found, an empty Passenger is returned.
func (p Passengers) Find(ssn string) Passenger {
	if one, ok := p[ssn]; ok {
		return one
	}
	return Passenger{}
}

// VisitUpdate calls visitor for each Passenger in the set. Updating their SSN's is not recommended.
func (p Passengers) VisitUpdate(visitor func(p *Passenger)) {
	for ssn, pp := range p {
		visitor(&pp)
		(p)[ssn] = pp
	}
}

// Manifest returns the SSN's of all Passengers in the set.
func (p Passengers) Manifest() []string {
	ssns := make([]string, 0, len(p))
	p.Visit(func(p Passenger) { ssns = append(ssns, p.SSN) })
	return ssns
}

// Charge prints a message that the Passenger has been charged "amount" money, and returns a copy with validTicket = true.
func (p Passenger) Charge(amount float64) Passenger {
	if p.HasValidTicket {
		return p // We already charged this Passenger.
	}
	fmt.Printf("Passenger with SSN %s: charged %.2f of arbitrary money\n", p.SSN, amount)
	p.HasValidTicket = true
	return p
}

// IsSenior returns true if the Passenger is a senior, and false otherwise.
// IsSenior detects age by extracting the last two digits from the SSN and treating them like an age.
func (p Passenger) IsSenior() bool {
	age, err := strconv.ParseInt(p.SSN[len(p.SSN)-2:], 10, 8)
	if err != nil {
		panic("invalid SSN: " + p.SSN)
	}
	return age >= SeniorAge
}
