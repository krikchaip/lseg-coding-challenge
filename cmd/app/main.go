package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/krikchaip/lseg-coding-challenge/internal/model"
)

const TIME_FORMAT = "15:04:05"

type TaskMonitor struct {
	pendingTasks map[string]*model.Task
}

func (this *TaskMonitor) Append(t model.Task) {
	pt, ok := this.pendingTasks[t.Description]

	if !ok {
		this.addNewTask(t)
		return
	}

	if !t.StartedAt.IsZero() {
		log.Printf("START record for task %q already existed.\n", t.Description)
		return
	}

	pt.EndedAt = t.EndedAt

	this.popTask(pt)
}

func (this *TaskMonitor) addNewTask(t model.Task) {
	if t.StartedAt.IsZero() {
		log.Printf("START record not found for task %q.\n", t.Description)
		return
	}

	if this.pendingTasks == nil {
		this.pendingTasks = map[string]*model.Task{t.Description: &t}
		return
	}

	this.pendingTasks[t.Description] = &t
}

func (this *TaskMonitor) popTask(t *model.Task) {
	if t.StartedAt.IsZero() {
		log.Printf("START record is missing for task %q.\n", t.Description)
		return
	}

	if t.EndedAt.IsZero() {
		log.Printf("END record is missing for task %q.\n", t.Description)
		return
	}

	minsTaken := t.EndedAt.Sub(t.StartedAt).Minutes()

	if 5 < minsTaken && minsTaken <= 10 {
		log.Printf(
			"[WARNING] task %q, pid %d took %.0fmins to complete.",
			t.Description,
			t.Pid,
			minsTaken,
		)
	} else if minsTaken > 10 {
		log.Printf(
			"[ERROR] task %q, pid %d took %.0fmins to complete.",
			t.Description,
			t.Pid,
			minsTaken,
		)
	}

	delete(this.pendingTasks, t.Description)
}

const LOGFILE_PATH = "data/logs.log"

func main() {
	var tm TaskMonitor

	file := readFile(LOGFILE_PATH)
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
