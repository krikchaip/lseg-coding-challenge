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
			log.Println(err)
			continue
		}

		var t Task

		t.Description = record[1]

		pid, err := strconv.Atoi(record[3])
		if err != nil {
			log.Println(err)
			continue
		}

		t.Pid = pid

		command := strings.TrimSpace(record[2])
		switch command {
		case "START":
			start, err := time.Parse(TIME_FORMAT, record[0])
			if err != nil {
				log.Println(err)
				continue
			}

			t.StartedAt = start
		case "END":
			end, err := time.Parse(TIME_FORMAT, record[0])
			if err != nil {
				log.Println(err)
				continue
			}

			t.EndedAt = end
		default:
			log.Println(fmt.Errorf("Unknown command %q", command))
			continue
		}

		fmt.Println(t)
	}
}
