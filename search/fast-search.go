package main

import "fmt"

// FastSearch implements SearchStrategy interface
type FastSearch struct{}

func (f *FastSearch) doSearch(filters map[string]int) {
	fmt.Printf("FAST search with following filters: %+v \n", filters)
}
