package main

import "fmt"

type Developer struct {
	Name string
	Age  int
}

func filterUnique(developers []Developer) []string {
	devMap := make(map[string]int)
	devNames := []string{}

	for _, dev := range developers {
		devMap[dev.Name] = devMap[dev.Name] + 1
	}

	for devName := range devMap {
		devNames = append(devNames, devName)
	}

	return devNames
}

func main() {
	fmt.Println("Filter Unique Challenge")
}
