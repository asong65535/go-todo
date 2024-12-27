package add

import (
	"encoding/csv"
	"os"
)

var ErrNoFirstLine = "First line not found"

/*
Purpose: Add to todo list (csv file)

	Cases: {
		- Add to empty list
		- Add to list that already exists
	}

	Outline: {
		- Gets passed a string from CLI
		- Opens and read file to see if new line exists
		- Go to end of file
		- Write new tasks, including ID and Datetime
		- Write buffer
		- Close file
		- Returns errors along the way if exists, otherwise exits with return nil
	}
*/
func Add(task string) error {

	// test records
	records := [][]string{
		{"ID", "Task", "Created"},
		{"1", "TEST", "00:00"},
	}

	categoryLine := [][]string{
		{"ID", "Task", "Created"},
	}

	// prints CWD
	// cwd, err := os.Getwd()
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(cwd)

	relativePath := "../../test/todo.csv"

	file, err := os.OpenFile(relativePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	csvWriter := csv.NewWriter(file)

	csvWriter.Write(categoryLine[0])

	for _, record := range records {
		err := csvWriter.Write(record)
		if err != nil {
			return err
		}
	}

	csvWriter.Flush()

	err = file.Close()
	if err != nil {
		return err
	}

	err = csvWriter.Error()
	if err != nil {
		return err
	}

	return nil
}
