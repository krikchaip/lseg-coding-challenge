package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const TIME_FORMAT = "15:04:05"

type Task struct {
	Description string
	Pid         int
	StartedAt   time.Time
	EndedAt     time.Time
}

func main() {
	file, err := os.Open("data/logs.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		var t Task

		t.Description = record[1]

		pid, err := strconv.Atoi(record[3])
		if err != nil {
			log.Fatal(err)
		}

		t.Pid = pid

		command := strings.TrimSpace(record[2])
		switch command {
		case "START":
			start, err := time.Parse(TIME_FORMAT, record[0])
			if err != nil {
				log.Fatal(err)
			}

			t.StartedAt = start
		case "END":
			end, err := time.Parse(TIME_FORMAT, record[0])
			if err != nil {
				log.Fatal(err)
			}

			t.EndedAt = end
		default:
			log.Fatal(fmt.Errorf("Unknown command %q", command))
		}

		fmt.Println(t)
	}
}
