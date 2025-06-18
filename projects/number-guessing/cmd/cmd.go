package cmd

import (
	"github.com/selahaddinislamoglu/golang-studies/projects/number-guessing/internal/handler"
	"github.com/spf13/cobra"
)

func NewCommand(handler handler.Handler) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "number-guessing",
		Short: "Play the number guessing game",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler.Run()
		},
	}

	return cmd
}
