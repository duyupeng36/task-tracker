package cmd

import (
	"github.com/spf13/cobra"
	"task-tracker/app"
)

func init() {
	deleteCmd.Flags().IntVarP(&id, "id", "i", 0, "ID of the task")
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a task by id",
	RunE: func(cmd *cobra.Command, args []string) error {
		if id == 0 {
			return cmd.Usage()
		}
		return app.App.DeleteTask(id)
	},
}
