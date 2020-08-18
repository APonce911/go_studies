package main

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	flightsQueue := Queue{
		Items: []Flight{
			Flight{
				Origin:      "CHINA",
				Destination: "MEXICO",
				Price:       500,
			},
		},
	}

	f := Flight{
		Origin:      "BRAZIL",
		Destination: "USA",
		Price:       200,
	}

	flightsQueue.Push(f)

	flightsQueueLength := len(flightsQueue.Items)
	if flightsQueueLength != 2 {
		t.Errorf("Expected 2 flights, got %v", flightsQueueLength)
	}

	lastFlight := flightsQueue.Items[flightsQueueLength-1]
	if lastFlight.Origin != "CHINA" {
		t.Errorf("Expected CHINA - MEXICO flight, got %v - %v", lastFlight.Origin, lastFlight.Destination)
	}
}

func TestPeak(t *testing.T) {
	flightsQueue := Queue{
		Items: []Flight{
			Flight{
				Origin:      "BRAZIL",
				Destination: "USA",
				Price:       200,
			},
			Flight{
				Origin:      "CHINA",
				Destination: "MEXICO",
				Price:       500,
			},
		},
	}

	flight, _ := flightsQueue.Peek()

	if flight.Origin != "CHINA" {
		fmt.Print(flightsQueue.Items)
		t.Errorf("Expected CHINA - MEXICO flight, got %v - %v", flight.Origin, flight.Destination)
	}
}

func TestPop(t *testing.T) {
	flightsQueue := Queue{
		Items: []Flight{
			Flight{
				Origin:      "BRAZIL",
				Destination: "USA",
				Price:       200,
			},
			Flight{
				Origin:      "CHINA",
				Destination: "MEXICO",
				Price:       500,
			},
		},
	}

	flight, _ := flightsQueue.Pop()

	if flight.Origin != "CHINA" {
		fmt.Print(flightsQueue.Items)
		t.Errorf("Expected to Pop CHINA - MEXICO flight, got %v - %v", flight.Origin, flight.Destination)
	}


	flightsQueueLength := len(flightsQueue.Items)
	if flightsQueueLength != 1 {
		t.Errorf("Expected 1 flight on queue, got %v", flightsQueueLength)
	}

}