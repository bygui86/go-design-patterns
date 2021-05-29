package main

type user struct {
	strategy Strategy
	filters  map[string]int
}

func initUser(strategy Strategy) *user {
	filters := make(map[string]int)
	return &user{
		strategy: strategy,
		filters:  filters,
	}
}

func (u *user) setStrategy(strategy Strategy) {
	u.strategy = strategy
}

func (u *user) getData() {
	u.strategy.doSearch(u.filters)
}
