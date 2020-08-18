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

	lastIndex := len(q.Items) - 1
	f := q.Items[lastIndex]

	q.Items = q.Items[:lastIndex]

	return f, nil
}

func (q *Queue) Push(f Flight) {
	q.Items = append([]Flight{f}, q.Items...)
}

func (q *Queue) Peek() (Flight, error) {
	// TODO Implement
	if q.IsEmpty() {
		return Flight{}, errors.New("StackIsEmpty")
	}

	return q.Items[len(q.Items)-1], nil
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
