package main

import (
	"io"
	"log"
	"os"
)

func main() {

	fileName := os.Args[len(os.Args)-1]

	file, err := os.Open("tasks/printfile/" + fileName)

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, file)
}
