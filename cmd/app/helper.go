package main

import (
	"encoding/csv"
	"io"
	"iter"
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

func readCSVLines(r io.Reader) iter.Seq[[]string] {
	cr := csv.NewReader(r)

	return func(yield func([]string) bool) {
		for {
			record, err := cr.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Println(err)
				continue
			}

			if !yield(record) {
				return
			}
		}
	}
}
