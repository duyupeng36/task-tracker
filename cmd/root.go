package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var id int
var status string
var description string

var rootCMD = &cobra.Command{
	Use:   "task-tracker",
	Short: "Task Tracker is a CLI tool for manager tasks",
	Long:  `Task Tracker is a CLI tool for manager tasks. It allows you to create, list, update and delete tasks`,
}

func init() {
	rootCMD.AddCommand(createCmd)
	rootCMD.AddCommand(deleteCmd)
	rootCMD.AddCommand(updateCmd)
	rootCMD.AddCommand(listCmd)

	rootCMD.PersistentFlags().StringVarP(&status, "status", "s", "", "Status of task")
	rootCMD.PersistentFlags().StringVarP(&description, "description", "d", "", "Description of task")
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
	}
}
