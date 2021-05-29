package main

import "fmt"

// Entrypoint
func main() {
	fmt.Println("Create categories")
	c1 := &category{"category 1"}
	c2 := &category{"category 2"}
	c3 := &category{"category 3"}

	fmt.Println("Create directories")
	d1 := &directory{
		name:     "Directory 1",
		children: []prototype{c1},
	}
	d2 := &directory{
		name:     "Directory 2",
		parents:  []directory{*d1},
		children: []prototype{d1, c2, c3},
	}

	fmt.Println("Open directory 2")
	d2.show("  ")

	fmt.Println("Clone and open directory 2")
	cloneFolder := d2.clone()
	cloneFolder.show("  ")
}
