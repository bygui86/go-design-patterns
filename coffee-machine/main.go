package main

import "fmt"

func main() {
	orderAnEspresso()

	orderADarkRoast()

	orderAnHouseBlend()
}

func orderAnEspresso() {
	// Order up an espresso, no condiments and print its description and cost
	espresso := &espresso{}
	fmt.Printf("%s $%.2f\n", espresso.description(), espresso.cost())
}

func orderADarkRoast() {
	// Make a DarkRoast object
	darkRoast := &darkRoast{}
	// Wrap it with a Mocha
	singleMocha := &mocha{
		beverage: darkRoast,
	}
	// Wrap it in a second Mocha
	doubleMocha := &mocha{
		beverage: singleMocha,
	}
	// Wrap it in a Whip.
	doubleMochaWhip := &whip{
		beverage: doubleMocha,
	}
	// Print its description and cost
	fmt.Printf("%s $%.2f\n", doubleMochaWhip.description(), doubleMochaWhip.cost())
}

func orderAnHouseBlend() {
	// Finally give us a HouseBlend with Soy, Mocha, and Whip.
	soyMochaWhipHouseBlend := &whip{
		beverage: &mocha{
			beverage: &soy{
				beverage: &houseBlend{},
			},
		},
	}
	fmt.Printf("%s $%.2f\n", soyMochaWhipHouseBlend.description(), soyMochaWhipHouseBlend.cost())
}
