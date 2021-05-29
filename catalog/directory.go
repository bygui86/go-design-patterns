package main

import "fmt"

// Directory type for node
type directory struct {
	name     string
	parents  []directory
	children []prototype
}

// Implement show function
func (f *directory) show(i string) {
	fmt.Println(i + f.name)
	for _, v := range f.children {
		v.show(i + i)
	}
}

// Implement clone function
func (f *directory) clone() prototype {
	cloneFolder := &directory{name: f.name + "_clone"}
	var tempChildren []prototype
	for _, i := range f.children {
		c := i.clone()
		tempChildren = append(tempChildren, c)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
