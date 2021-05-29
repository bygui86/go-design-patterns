package control

type SimpleRemoteControl struct {
	slot Command
}

/*
	We have a method for setting the Command the slot is going to control. This could be called
	multiple times if the client of this code wanted to change the behaviour of the remote button.
*/
func (r *SimpleRemoteControl) SetCommand(command Command) {
	r.slot = command
}

/*
	This method is called when the button is pressed. All we do is take the current Command bound to the
	slot and call its Execute() method.
*/
func (r *SimpleRemoteControl) PressButton() {
	r.slot.Execute()
}
