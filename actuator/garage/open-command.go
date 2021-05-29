package garage

// GarageDoorOpenCommand implements control.Command interface
type GarageDoorOpenCommand struct {
	garage *Garage
}

func NewGarageDoorOpenCommand(g *Garage) *GarageDoorOpenCommand {
	return &GarageDoorOpenCommand{
		garage: g,
	}
}

// Execute results in a combo of different sequential commands
func (g *GarageDoorOpenCommand) Execute() {
	g.garage.doorUp()
	g.garage.lightOn()
}
