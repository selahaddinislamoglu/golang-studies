package cmd

import (
	"os"

	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/internal/handler"
	"github.com/spf13/cobra"
)

func NewAddCmd(handler handler.Handler) *cobra.Command {
	return &cobra.Command{
		Use:   "add [description]",
		Short: "Add a new task",
		Long:  `Add a new task with the specified description. The task will be created with the current date and time as its creation time.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) < 1 {
				cmd.Println("Please provide a task description.")
				os.Exit(1)
			}

			description := args[0]
			id, err := handler.AddTask(description)
			if err != nil {
				cmd.Printf("Error adding task: %v\n", err)
				os.Exit(1)
			}

			cmd.Printf("Task added successfully. ID:%d\n", id)
		},
	}
}
