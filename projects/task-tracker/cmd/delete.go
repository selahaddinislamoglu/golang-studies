package cmd

import (
	"fmt"
	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/internal/handler"
	"github.com/spf13/cobra"
	"os"
)

func NewDeleteCmd(handler handler.Handler) *cobra.Command {
	return &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a task",
		Long:  `Delete a task by its ID. The task will be removed from the task list and the data file.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) < 1 {
				fmt.Println("Please provide a task id.")
				os.Exit(1)
			}

			id := args[0]
			err := handler.DeleteTask(id)
			if err != nil {
				fmt.Printf("Error adding task: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Task deleted successfully. ID:%s\n", id)
		},
	}
}
