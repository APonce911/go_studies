package main

import "testing"

func TestPush(t *testing.T) {
	flightsStack := Stack{}
	f := Flight{
		Origin:      "US",
		Destination: "Europe",
		Price:       100}

	flightsStack.Push(f)

	if len(flightsStack.Flights) != 1 {
		t.Errorf("Expected 1 flight, got %v", len(flightsStack.Flights))
	}
}

func TestPeak(t *testing.T) {
	flightsStack := Stack{
		Flights: []Flight{
			Flight{
				Origin:      "US",
				Destination: "Europe",
				Price:       100,
			},
		},
	}

	flight, _ := flightsStack.Peek()

	if flight.Origin != "US" {
		t.Errorf("Expected US origin, got %v", flight.Origin)
	}
}

func TestPop(t *testing.T) {
	flightsStack := Stack{
		Flights: []Flight{
			Flight{
				Origin:      "US",
				Destination: "Europe",
				Price:       100,
			},
		},
	}

	flight, _ := flightsStack.Pop()

	if len(flightsStack.Flights) != 0 {
		t.Errorf("Expected empty Stack, got %v flights", len(flightsStack.Flights))
	}

	if flight.Origin != "US" {
		t.Errorf("Expected popped flight with US origin, got %v", flight.Origin)
	}
}
