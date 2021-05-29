package main

import "fmt"

// Category type for node
type category struct {
	name string
}

// Implement show function
func (f *category) show(i string) {
	fmt.Println(i + f.name)
}

// Implement clone function
func (f *category) clone() prototype {
	return &category{name: f.name + "_clone"}
}
