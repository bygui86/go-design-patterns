package main

import "fmt"

// FastSearch implements Strategy interface
type FastSearch struct{}

func (f *FastSearch) doSearch(filters map[string]int) {
	fmt.Printf("FAST search with following filters: %+v \n", filters)
}
