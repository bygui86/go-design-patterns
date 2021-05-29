package main

// transport implements iTransport interface
type transport struct {
	name  string
	speed uint
}

func (t *transport) SetName(n string) {
	t.name = n
}

func (t *transport) GetName() string {
	return t.name
}

func (t *transport) SetSpeed(s uint) {
	t.speed = s
}

func (t *transport) GetSpeed() uint {
	return t.speed
}
