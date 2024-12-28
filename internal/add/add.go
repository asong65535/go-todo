package add

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"time"
)

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

	categoryLine := []string{"ID", "Description", "CreatedAt", "IsComplete"}

	relativePath := "../../test/todo.csv"

	isFileExists := checkFileExists(relativePath)

	var file *os.File
	var fileError error

	openFileFlag := os.O_RDWR

	// if csv exists, open it. otherwise create one
	if isFileExists {
		file, fileError = os.OpenFile(relativePath, openFileFlag, 0644)
		if fileError != nil {
			return fileError
		}
	} else {
		file, fileError = os.Create(relativePath)
		if fileError != nil {
			return fileError
		}
	}

	oldList, err := csv.NewReader(file).ReadAll() // initial todo list
	if err != nil {
		return err
	}

	taskID := len(oldList)

	csvWriter := csv.NewWriter(file)

	if !isFileExists {
		taskID = 1 // manually set since new file would set it to 0
		csvWriter.Write(categoryLine)
	}

	parsedTask := taskParser(taskID, task)

	csvWriter.Write(parsedTask)

	csvWriter.Flush()

	err = csvWriter.Error()
	if err != nil {
		return err
	}

	return nil
}

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	return !errors.Is(err, os.ErrNotExist)
}

func taskParser(id int, task string) []string {
	var parsedString []string

	parsedString = append(parsedString, strconv.Itoa(id))
	parsedString = append(parsedString, task)
	parsedString = append(parsedString, time.Now().Format(time.DateTime))
	parsedString = append(parsedString, "false")

	return parsedString
}
