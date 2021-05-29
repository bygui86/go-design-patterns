package control

import (
	"fmt"
	"reflect"
	"strings"
)

type MultiRemoteControl struct {
	onCommands  []Command
	offCommands []Command
}

func NewMultiRemoteControl(cmdNum int) *MultiRemoteControl {
	return &MultiRemoteControl{
		onCommands:  make([]Command, cmdNum),
		offCommands: make([]Command, cmdNum),
	}
}

/*
	The setCommand() method takes a slot position and an On and Off command to be stored in that slot.
	It puts these commands in the on and off arrays for later use.
*/
func (r *MultiRemoteControl) SetCommand(slot int, onCommand, offCommand Command) {
	r.onCommands[slot] = onCommand
	r.offCommands[slot] = offCommand
}

// When an On or Off button is pressed, the hardware takes care of calling the corresponding methods.
func (r *MultiRemoteControl) PressOnButton(slot int) {
	r.onCommands[slot].Execute()
}

func (r *MultiRemoteControl) PressOffButton(slot int) {
	r.offCommands[slot].Execute()
}

// Implementing String() to print out each slot and its corresponding command.
func (r *MultiRemoteControl) String() string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString("------ Remote Control -------\n")
	for i := range r.onCommands {
		if r.onCommands[i] == nil {
			continue
		}
		onClass := r.getClassName(r.onCommands[i])
		offClass := r.getClassName(r.offCommands[i])
		strBuilder.WriteString(fmt.Sprintf("[slot %d] %s - %s\n", i, onClass, offClass))
	}
	strBuilder.WriteString("-----------------------------")
	return strBuilder.String()
}

func (r *MultiRemoteControl) getClassName(myVar interface{}) string {
	if t := reflect.TypeOf(myVar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
