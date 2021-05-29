package singleton

import "sync"

// declare variable
var doOnce sync.Once
var instance *singleton = nil

// declaration defined type
type singleton struct {
	title string
	sync.RWMutex
}

// defined type with interface
type Singleton interface {
	SetTitle(t string)
	GetTitle() string
}

// Get only one object
func GetInstance() Singleton {
	doOnce.Do(
		func() {
			instance = new(singleton)
		},
	)
	return instance
}

// Setter for singleton variable
func (s *singleton) SetTitle(t string) {
	s.Lock()
	defer s.Unlock()
	s.title = t
}

// Getter singleton variable
func (s *singleton) GetTitle() string {
	s.RLock()
	defer s.RUnlock()
	return s.title
}
