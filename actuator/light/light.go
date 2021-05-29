package light

import "fmt"

type Light struct {
	Name string
}

func (l *Light) on() {
	fmt.Printf("%s on \n", l.Name)
}

func (l *Light) off() {
	fmt.Printf("%s off \n", l.Name)
}
