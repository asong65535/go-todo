package list

import (
	//"encoding/csv"
	//"errors"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	//"github.com/mergestat/timediff"
)

/*
Notes:
- replace commas with "\t"?
- get flag [-a]
- loop through csv and parse with fmt.Fprintln
*/
func List() error {

	// demo written with hardcoded values

	/*
		w := tabwriter.NewWriter(os.Stdout, 0, 4, 4, ' ', tabwriter.Debug)

		// test list
		fmt.Fprintln(w, "ID\tTask\tCreated")
		fmt.Fprintln(w, "1\tTidy up my desk\ta minute ago")
		fmt.Fprintln(w, "3\tChange my keyboard mapping to use escape/control\ta few seconds ago")

		// \n
		fmt.Fprint(w, "\n")

		// test list -a
		fmt.Fprintln(w, "ID\tTask\tCreated\tDone")
		fmt.Fprintln(w, "1\tTidy up my desk\t2 minutes ago\tfalse")
		fmt.Fprintln(w, "2\tWrite up documentation for new project feature\ta minute ago\ttrue")
		fmt.Fprintln(w, "3\tChange my keyboard mapping to use escape/control\ta minute ago\tfalse")

		w.Flush()
	*/

	w := tabwriter.NewWriter(os.Stdout, 0, 4, 4, ' ', 0)

	// TODO: change to "./test/todo.csv for prod"
	filepath := "../../../test/todo.csv"

	openFileFlag := os.O_RDWR

	//var allFlag bool = false
	var allFlag bool = true

	file, fileError := os.OpenFile(filepath, openFileFlag, 0644)
	if fileError != nil {
		return fileError
	}

	csvReader, readError := csv.NewReader(file).ReadAll()
	if readError != nil {
		return readError
	}

	if !allFlag {
		for i, task := range csvReader {

			fmt.Fprintln(w, taskParser(task, i, allFlag))
		}
	} else {
		for i, task := range csvReader {

			fmt.Fprintln(w, taskParser(task, i, allFlag))
		}
	}

	flushError := w.Flush()
	if flushError != nil {
		return flushError
	}

	return nil
}

func taskParser(task []string, line int, flag bool) string {
	parsedTask := ""

	b, err := strconv.ParseBool(task[3])
	if err != nil && line != 0 {
		fmt.Println(err)
	}

	if !b {
		return parsedTask
	}

	if !flag {
		if line == 0 {
			for i, word := range task[:3] {
				switch i {
				case 1:
					parsedTask += "Task\t"
				case 2:
					parsedTask += "Created\t"
				default:
					parsedTask += word + "\t"
				}
			}
			return parsedTask
		}

		for _, word := range task[:3] {
			parsedTask += word + "\t"
		}

	} else {
		if line == 0 {
			for i, word := range task {
				switch i {
				case 1:
					parsedTask += "Task\t"
				case 2:
					parsedTask += "Created\t"
				case 3:
					parsedTask += "Done"
				default:
					parsedTask += word + "\t"
				}
			}
			return parsedTask
		}

		for _, word := range task {
			parsedTask += word + "\t"
		}

	}

	return parsedTask
}
