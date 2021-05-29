package main

// scooter inherit from transport
type scooter struct {
	transport
}

func newScooter() iTransport {
	return &scooter{
		transport: transport{
			name:  "Scooter",
			speed: 4,
		},
	}
}
