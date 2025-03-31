package main

import (
	"log"
	"strings"

	"github.com/krikchaip/lseg-coding-challenge/internal/core"
	"github.com/krikchaip/lseg-coding-challenge/internal/model"
)

const (
	LOGFILE_PATH = "data/logs.log"
)

func main() {
	var tm core.TaskMonitor

	file := readFile(LOGFILE_PATH)
	defer file.Close()

	for record := range readCSVLines(file) {
		timestamp, description, entry, pid := record[0], record[1], record[2], record[3]

		tl, err := model.NewTaskLogFromStrings(
			timestamp,
			description,
			strings.TrimSpace(entry),
			pid,
		)
		// ignore the current line if it is somehow incorrect
		if err != nil {
			log.Println(err)
			continue
		}

		if err := tm.AppendLog(tl); err != nil {
			log.Println(err)
		}
	}
}
