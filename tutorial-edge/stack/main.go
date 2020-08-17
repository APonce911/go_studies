package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Flights []Flight
}

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func (s *Stack) Pop() (Flight, error) {
	if s.IsEmpty() {
		return Flight{}, errors.New("StackIsEmpty")
	}

	lastIndex := len(s.Flights) - 1
	f := s.Flights[lastIndex]

	s.Flights = s.Flights[:lastIndex]

	return f, nil
}

func (s *Stack) Push(f Flight) {
	s.Flights = append(s.Flights, f)
}

func (s *Stack) Peek() (Flight, error) {
	if s.IsEmpty() {
		return Flight{}, errors.New("StackIsEmpty")
	}

	return s.Flights[len(s.Flights)-1], nil
}

func (s *Stack) IsEmpty() bool {
	if len(s.Flights) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Go Stack Implementation")
}
