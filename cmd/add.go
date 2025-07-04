/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"tasks/internal/add"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use: "add",
	// Short: "A brief description of your command",
	Short: "Add to Todo",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var task string

		// args returns words in separate elems in []
		// this loop concats to a single str [task] which is passed to Add()
		for i, word := range args {
			task += word
			if i == len(args)-1 {
				break
			}
			task += " "
		}

		// Error checking
		err := add.Add(task)
		if err != nil {
			log.Fatal("Unable to add to database: ", err)
		} else {
			fmt.Printf("Task successfully added!")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
