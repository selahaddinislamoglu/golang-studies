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

func NewUpdateStatusInProgressCmd(handler handler.Handler) *cobra.Command {
	return &cobra.Command{
		Use:   "mark-in-progress [id]",
		Short: "Mark a task as in progress",
		Long:  `Mark a task as in progress by its ID. The task status will be updated in the task list and the data file.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]

			err := handler.UpdateTaskStatusInProgress(id)
			if err != nil {
				cmd.PrintErrf("Error marking task as in progress: %v\n", err)
				return
			}

			cmd.Printf("Task marked as in progress successfully. ID:%s\n", id)
		},
	}
}

func NewUpdateStatusDoneCmd(handler handler.Handler) *cobra.Command {
	return &cobra.Command{
		Use:   "mark-done [id]",
		Short: "Mark a task as done",
		Long:  `Mark a task as done by its ID. The task status will be updated in the task list and the data file.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]

			err := handler.UpdateTaskStatusDone(id)
			if err != nil {
				cmd.PrintErrf("Error marking task as done: %v\n", err)
				return
			}

			cmd.Printf("Task marked as done successfully. ID:%s\n", id)
		},
	}
}

func NewUpdateStatusToDoCmd(handler handler.Handler) *cobra.Command {
	return &cobra.Command{
		Use:   "mark-todo [id]",
		Short: "Mark a task as todo",
		Long:  `Mark a task as todo by its ID. The task status will be updated in the task list and the data file.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]

			err := handler.UpdateTaskStatusTodo(id)
			if err != nil {
				cmd.PrintErrf("Error marking task as todo: %v\n", err)
				return
			}

			cmd.Printf("Task marked as todo successfully. ID:%s\n", id)
		},
	}
}
