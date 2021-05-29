package main

func main() {
	// run fast search
	fast := &FastSearch{}
	user1 := initUser(fast)
	user1.filters["role"] = 1
	user1.getData()

	// run deep search
	deep := &DeepSearch{}
	user2 := initUser(deep)
	user2.filters["role"] = 1
	user2.getData()
}
