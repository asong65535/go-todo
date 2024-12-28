package list

import (
	//"errors"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

// TODO: pass in [-a] flag
func List(flag bool) error {

	w := tabwriter.NewWriter(os.Stdout, 0, 4, 4, ' ', 0)

	// TODO: change to "./test/todo.csv for prod"
	//filepath := "../../../test/todo.csv"
	filepath := "./test/todo.csv"

	openFileFlag := os.O_RDWR

	// TODO: remove test flags
	//var allFlag bool = false
	//var allFlag bool = true

	categoryRow := "ID\tTask\tCreated"
	underlineRow := "--\t----\t-------"
	if flag {
		categoryRow = "ID\tTask\tCreated\tDone"
		underlineRow = "--\t----\t------\t----"
	}

	file, fileError := os.OpenFile(filepath, openFileFlag, 0644) // -rw-r--r--
	if fileError != nil {
		return fileError
	}

	csvReader, readError := csv.NewReader(file).ReadAll()
	if readError != nil {
		return readError
	}

	// setup table
	fmt.Fprintln(w, categoryRow)
	fmt.Fprintln(w, underlineRow)

	// push rest of table to writer
	for _, task := range csvReader[1:] {

		isComplete, err := strconv.ParseBool(task[3])
		if err != nil {
			return nil
		}

		if isComplete && !flag {
			continue
		}

		fmt.Fprintln(w, taskParser(task, flag))
	}

	// write to terminal
	flushError := w.Flush()
	if flushError != nil {
		return flushError
	}

	return nil
}

func taskParser(task []string, allFlag bool) string {

	parsedTask := ""

	if !allFlag {
		for i, word := range task[:3] {
			if i == 2 {
				parsedTask += TimeParser(task[i]) + "\t"
				continue
			}
			parsedTask += word + "\t"
		}

	} else {
		for i, word := range task {
			if i == 2 {
				parsedTask += TimeParser(task[i]) + "\t"
				continue
			}
			parsedTask += word + "\t"
		}
	}

	return parsedTask
}

func TimeParser(taskTime string) string {

	createdTime, err := time.Parse(time.DateTime, taskTime)
	if err != nil {
		log.Fatal(err)
	}

	return timediff.TimeDiff(createdTime.Local())
}
