package garage

import "fmt"

type Garage struct{}

func (l *Garage) doorUp() {
	fmt.Println("door up")
}

func (l *Garage) doorDown() {
	fmt.Println("door down")
}

func (l *Garage) lightOn() {
	fmt.Println("light on")
}

func (l *Garage) lightOff() {
	fmt.Println("light off")
}
