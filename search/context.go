package main

type user struct {
	strategy SearchStrategy
	filters  map[string]int
}

func initUser(strategy SearchStrategy) *user {
	filters := make(map[string]int)
	return &user{
		strategy: strategy,
		filters:  filters,
	}
}

func (u *user) setStrategy(strategy SearchStrategy) {
	u.strategy = strategy
}

func (u *user) getData() {
	u.strategy.doSearch(u.filters)
}
