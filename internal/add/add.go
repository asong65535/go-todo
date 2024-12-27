package add

import (
	"encoding/csv"
	"fmt"
	"os"
)

// import "fmt"
func Add(task string) error {

	records := [][]string{
		{"ID", "Task", "Created"},
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Println(cwd)

	relativePath := "../../test/todo.csv"

	file, err := os.OpenFile(relativePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	w := csv.NewWriter(file)

	for _, record := range records {
		err := w.Write(record)
		if err != nil {
			return err
		}
	}

	w.Flush()

	err = file.Close()
	if err != nil {
		return err
	}

	err = w.Error()
	if err != nil {
		return err
	}

	return nil
}
