package complete

import (
	"encoding/csv"
	"os"
	"strconv"
)

/*
	NOTES:
	Get taskID from cobra
	Look through csv to find ID

*/

func Complete(id int) error {

	pathToList := "./test/todo.csv"
	pathToTemp := "./test/todo_temp.csv"

	// load original list to mem
	original, err := os.Open(pathToList)
	if err != nil {
		return err
	}
	defer original.Close()

	rows, err := csv.NewReader(original).ReadAll()
	if err != nil {
		return err
	}

	// create temp file
	temp, err := os.Create(pathToTemp)
	if err != nil {
		return err
	}
	defer temp.Close()

	csvWriter := csv.NewWriter(temp)
	defer csvWriter.Flush()

	csvWriter.Write(rows[0]) // write category line

	for _, row := range rows[1:] {
		currentID, err := strconv.Atoi(row[0])
		if err != nil {
			return err
		}

		if id == currentID {
			row[3] = "true"
			csvWriter.Write(row)
			continue
		}
		csvWriter.Write(row)
	}

	if err := csvWriter.Error(); err != nil {
		return err
	}

	// replace original list with new list
	if err := os.Rename(pathToTemp, pathToList); err != nil {
		return err
	}

	return nil
}
