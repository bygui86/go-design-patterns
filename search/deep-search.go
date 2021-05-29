package main

import "fmt"

// DeepSearch implements Strategy interface
type DeepSearch struct{}

func (f *DeepSearch) doSearch(filters map[string]int) {
	fmt.Printf("DEEP search with following filters: %+v \n", filters)
}
