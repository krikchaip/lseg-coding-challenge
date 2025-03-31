package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/krikchaip/lseg-coding-challenge/internal/core"
	"github.com/krikchaip/lseg-coding-challenge/internal/model"
)

const (
	TIME_FORMAT  = "15:04:05"
	LOGFILE_PATH = "data/logs.log"
)

func main() {
	var tm core.TaskMonitor

	file := readFile(LOGFILE_PATH)
	defer file.Close()

	for record := range readCSVLines(file) {
		var t model.Task

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

		tm.Append(t)
	}
}
