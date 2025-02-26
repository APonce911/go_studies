package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Items []Flight
}

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func (q *Queue) Pop() (Flight, error) {
	if q.IsEmpty() {
		return Flight{}, errors.New("StackIsEmpty")
	}

	f := q.Items[0]
	q.Items = q.Items[1:]

	return f, nil
}

func (q *Queue) Push(f Flight) {
	q.Items = append(q.Items, f)
}

func (q *Queue) Peek() (Flight, error) {
	// TODO Implement
	if q.IsEmpty() {
		return Flight{}, errors.New("StackIsEmpty")
	}

	return q.Items[0], nil
}

func (q *Queue) IsEmpty() bool {
	if len(q.Items) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Go Queue Implementation")
}
