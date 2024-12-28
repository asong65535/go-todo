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

/*
Notes:
- replace commas with "\t"?
- get flag [-a]
- loop through csv and parse with fmt.Fprintln
*/

func List() error {

	w := tabwriter.NewWriter(os.Stdout, 0, 4, 4, ' ', 0)

	// TODO: change to "./test/todo.csv for prod"
	filepath := "../../../test/todo.csv"

	openFileFlag := os.O_RDWR

	//var allFlag bool = false
	var allFlag bool = true

	file, fileError := os.OpenFile(filepath, openFileFlag, 0644) // -rw-r--r--
	if fileError != nil {
		return fileError
	}

	csvReader, readError := csv.NewReader(file).ReadAll()
	if readError != nil {
		return readError
	}

	categoryRow := "ID\tTask\tCreated"
	if allFlag {
		categoryRow = "ID\tTask\tCreated\tDone"
	}

	fmt.Fprintln(w, categoryRow)

	for _, task := range csvReader[1:] {

		isComplete, err := strconv.ParseBool(task[3])
		if err != nil {
			return nil
		}

		if isComplete {
			continue
		}

		fmt.Fprintln(w, taskParser(task, allFlag))
	}

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
				parsedTask += TimeParser(task[i])
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
