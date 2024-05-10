package cmd

import (
	"github.com/spf13/cobra"
	"lexilift/app"
	"log/slog"
)

var (
	rootCmd = &cobra.Command{
		Use:   "lexilift",
		Short: "learn english words with flash card",
		Run: func(cmd *cobra.Command, args []string) {
			if err := app.Run(false); err != nil {
				slog.Error(err.Error())
			}
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
