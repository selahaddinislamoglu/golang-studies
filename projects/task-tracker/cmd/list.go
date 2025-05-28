package cmd

import (
	"fmt"

	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/internal/handler"
	"github.com/spf13/cobra"
)

func NewListCmd(handler handler.Handler) *cobra.Command {
	return &cobra.Command{
		Use:   "list [options]",
		Short: "List all tasks",
		Long:  `List all tasks in the task list.`,
		Args:  cobra.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {

			var tasks [][]byte
			var err error
			if len(args) == 0 {
				tasks, err = handler.ListTasks()
				if err != nil {
					cmd.PrintErrf("Error listing tasks: %v\n", err)
					return
				}
			} else if len(args) == 1 {
				status := args[0]
				tasks, err = handler.ListTasksByStatus(status)
				if err != nil {
					cmd.PrintErrf("Error listing tasks by status '%s': %v\n", status, err)
					return
				}
			} else {
				cmd.PrintErr("Invalid number of arguments. Use 'list' to list all tasks or 'list [status]' to list tasks by status.")
				return
			}
			if len(tasks) == 0 {
				fmt.Println("No tasks found.")
				return
			}

			for _, task := range tasks {
				fmt.Println(string(task))
			}
		},
	}
}
