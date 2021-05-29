package garage

// GarageDoorCloseCommand implements control.Command interface
type GarageDoorCloseCommand struct {
	garage *Garage
}

func NewGarageDoorCloseCommand(g *Garage) *GarageDoorCloseCommand {
	return &GarageDoorCloseCommand{
		garage: g,
	}
}

// Execute results in a combo of different sequential commands
func (g *GarageDoorCloseCommand) Execute() {
	g.garage.doorDown()
	g.garage.lightOff()
}
