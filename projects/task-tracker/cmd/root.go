package cmd

import (
	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/internal/handler"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "tast-tracker [command]",
	Short:   "tast-tracker is a simple task management CLI",
	Long:    `tast-tracker is a simple task management CLI application that allows you to create, view, and manage tasks from the command line.`,
	Version: "0.1.0",
	Example: `tast-tracker add "My new task"`,
	Args:    cobra.MinimumNArgs(1),
}

func NewRootCmd(handler handler.Handler) *cobra.Command {
	rootCmd.AddCommand(NewAddCmd(handler))
	rootCmd.AddCommand(NewDeleteCmd(handler))
	rootCmd.AddCommand(NewUpdateDescriptionCmd(handler))
	rootCmd.AddCommand(NewUpdateStatusInProgressCmd(handler))
	rootCmd.AddCommand(NewUpdateStatusDoneCmd(handler))
	rootCmd.AddCommand(NewUpdateStatusToDoCmd(handler))
	rootCmd.AddCommand(NewListCmd(handler))
	return rootCmd
}
