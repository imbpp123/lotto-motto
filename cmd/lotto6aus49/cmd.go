package lotto6aus49

import "github.com/spf13/cobra"

var (
	showCmd = &cobra.Command{
		Use:   "6aus49",
		Short: "Show last N rows in results",
		RunE:  runShowCommand,
	}
)

func RegisterNewCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(showCmd)
}
