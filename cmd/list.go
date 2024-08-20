package cmd

import (
	"github.com/spf13/cobra"
	"task-tracker/app"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list tasks",
	Run: func(cmd *cobra.Command, args []string) {
		app.App.DisplayTasks(status)
	},
}
