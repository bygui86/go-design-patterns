package main

import (
	"fmt"
	"time"
)

type Server struct {
	port           string
	timeout        time.Duration
	maxConnections int
}

type Option func(*Server)

func main() {
	server1 := NewServer("4000", 50*time.Second, 500)

	server2 := NewServerWithOptions("4000")
	server3 := NewServerWithOptions("4000", WithTimeout(50*time.Second))
	server4 := NewServerWithOptions("4000", WithMaxConnections(500))
	server5 := NewServerWithOptions("4000", WithTimeout(50*time.Second), WithMaxConnections(500))

	fmt.Printf("server1: %+v", server1)
	fmt.Printf("server2: %+v", server2)
	fmt.Printf("server3: %+v", server3)
	fmt.Printf("server4: %+v", server4)
	fmt.Printf("server5: %+v", server5)
}

// NewServer - Classic constructor
func NewServer(port string, timeout time.Duration, maxConnections int) *Server {
	return &Server{
		port:           port,
		timeout:        timeout,
		maxConnections: maxConnections,
	}
}

// NewServerWithOptions - Flexbile contructor taking options as input
func NewServerWithOptions(port string, options ...Option) *Server {
	server := &Server{
		port: port,
	}

	// looping over options and calling them
	// let them do their job
	for _, option := range options {
		option(server)
	}

	return server
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConnections(maxConnecetions int) Option {
	return func(s *Server) {
		s.maxConnections = maxConnecetions
	}
}
