package main

import (
	"fmt"
	"os"
	"tasks/internal/list"
)

func main() {
	if err := list.List(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}
