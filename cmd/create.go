package cmd

import (
	"github.com/spf13/cobra"
	"task-tracker/app"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new task",
	RunE: func(cmd *cobra.Command, args []string) error {

		if status == "" {
			status = "todo"
		}

		if description == "" {
			return cmd.Usage()
		}

		return app.App.AddTask(status, description)
	},
}
