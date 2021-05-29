package main

type Strategy interface {
	doSearch(filters map[string]int)
}
