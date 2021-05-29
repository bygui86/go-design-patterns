package light

// LightOffCommand implements control.Command interface
type LightOffCommand struct {
	light *Light
}

/*
	The constructor is passed the specific Light that this Command is going to control - say the living room Light
	- and stashes it in the Light instance variable. When execute gets called, this is the Light object that is
	going to be the Receiver of the request.
*/
func NewLightOffCommand(l *Light) *LightOffCommand {
	return &LightOffCommand{
		light: l,
	}
}

// Execute calls the off() method on the receiving object, which is the Light we are controlling
func (l *LightOffCommand) Execute() {
	l.light.off()
}
