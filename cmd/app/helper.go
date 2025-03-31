package main

import (
	"log"
	"os"
)

func readFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
