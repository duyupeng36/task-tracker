package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"task-tracker/app"
)

func init() {
	updateCmd.Flags().IntVarP(&id, "id", "i", 0, "ID of the task")
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update task status or description",
	RunE: func(cmd *cobra.Command, args []string) error {
		if id == 0 {
			return errors.New("id is required")
		}
		if status == "" && description == "" {
			return errors.New("you must set either --status or --description")
		}

		return app.App.UpdateTask(id, status, description)
	},
}
