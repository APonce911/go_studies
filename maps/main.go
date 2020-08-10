package main

import "fmt"

func main() {
	myInts := make(map[int]int)
	myInts[1]=10
	myInts[2]=20
	myInts[1]=30

	fmt.Println(myInts)
}