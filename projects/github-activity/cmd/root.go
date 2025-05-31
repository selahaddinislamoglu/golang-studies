package cmd

import (
	"github.com/selahaddinislamoglu/golang-studies/projects/github-activity/internal/handler"
	"github.com/spf13/cobra"
)

func NewRootCmd(handler handler.HandlerInterface) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:     "github-activity <username>",
		Short:   "github-activity is a CLI tool to fetch and display GitHub user activity",
		Long:    `github-activity is a command-line application that allows you to fetch and display the activity of a GitHub user.`,
		Version: "0.1.0",
		Example: `github-activity selahaddinislamoglu`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cmd.Println("Please provide a GitHub username.")
				return
			}
			username := args[0]
			activities, err := handler.GetUserActivity(username)
			if err != nil {
				cmd.Println("Error fetching user activity:", err)
				return
			}

			if len(activities) == 0 {
				cmd.Println("No activity found for user:", username)
				return
			}

			for _, event := range activities {
				cmd.Println(event)
			}
		},
	}

	return rootCmd
}
