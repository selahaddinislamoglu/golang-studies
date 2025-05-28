package cmd

import (
	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/internal/handler"
	"github.com/spf13/cobra"
)

func NewUpdateDescriptionCmd(handler handler.Handler) *cobra.Command {
	return &cobra.Command{
		Use:   "update [id] [description]",
		Short: "Update a task",
		Long:  `Update a task by its ID and description. The task will be updated in the task list and the data file.`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			description := args[1]

			err := handler.UpdateTaskDescription(id, description)
			if err != nil {
				cmd.PrintErrf("Error updating task: %v\n", err)
				return
			}

			cmd.Printf("Task updated successfully. ID:%s, Description:%s\n", id, description)
		},
	}
}
