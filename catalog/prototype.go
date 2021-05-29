package main

// Implement interface for prototype pattern
type prototype interface {
	show(string)
	clone() prototype
}
