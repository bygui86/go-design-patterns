package main

import (
	"fmt"
	"time"

	"github.com/bygui86/go-design-patterns/actuator/control"
	"github.com/bygui86/go-design-patterns/actuator/garage"
	"github.com/bygui86/go-design-patterns/actuator/light"
)

func main() {
	simpleExample()

	fmt.Println()
	fmt.Println()

	comboExample()

	fmt.Println()
	fmt.Println()

	multiExample()
}

func simpleExample() {
	fmt.Println("SIMPLE CONTROL")

	// We create a Light object, this will be the Receiver of the request
	bulb := &light.Light{Name: "bulb"}

	// The remote is our Invoker; it will be passed a command object that can be used to make requests
	remoteControl := &control.SimpleRemoteControl{}

	// Create a command and pass the Receiver to it
	onCmd := light.NewLightOnCommand(bulb)

	// Pass the command to the Invoker
	remoteControl.SetCommand(onCmd)

	// We simulate the button being pressed
	remoteControl.PressButton()

	time.Sleep(3 * time.Second)

	// Pass the new command to the invoker
	remoteControl.SetCommand(light.NewLightOffCommand(bulb))
	remoteControl.PressButton()
}

func comboExample() {
	fmt.Println("COMBO CONTROL")

	myGarage := &garage.Garage{}

	remoteControl := &control.SimpleRemoteControl{}
	remoteControl.SetCommand(garage.NewGarageDoorOpenCommand(myGarage))
	remoteControl.PressButton()

	time.Sleep(3 * time.Second)

	remoteControl.SetCommand(garage.NewGarageDoorCloseCommand(myGarage))
	remoteControl.PressButton()
}

func multiExample() {
	fmt.Println("MULTI CONTROL")
	fmt.Println()

	livingBulb := &light.Light{Name: "living room"}
	bedroomBulb := &light.Light{Name: "bedroom"}
	bathroomBulb := &light.Light{Name: "bathroom"}

	remoteControl := control.NewMultiRemoteControl(3)
	remoteControl.SetCommand(
		0,
		light.NewLightOnCommand(livingBulb),
		light.NewLightOffCommand(livingBulb),
	)
	remoteControl.SetCommand(
		1,
		light.NewLightOnCommand(bedroomBulb),
		light.NewLightOffCommand(bedroomBulb),
	)
	remoteControl.SetCommand(
		2,
		light.NewLightOnCommand(bathroomBulb),
		light.NewLightOffCommand(bathroomBulb),
	)

	fmt.Println(remoteControl.String())
	fmt.Println()

	remoteControl.PressOnButton(0)
	time.Sleep(1 * time.Second)

	remoteControl.PressOnButton(1)
	time.Sleep(1 * time.Second)

	remoteControl.PressOffButton(2)
	time.Sleep(1 * time.Second)

	remoteControl.PressOffButton(1)
	time.Sleep(1 * time.Second)
}
