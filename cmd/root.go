package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "netflux",
	Short: "Netflux is the CLI to run the Netflux API server.",
	Long:  "Netflux is the CLI to run the Netflux API server.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}
