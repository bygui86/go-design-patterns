package main

type SearchStrategy interface {
	doSearch(filters map[string]int)
}
