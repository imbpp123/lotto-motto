package eurojackpot

import "github.com/spf13/cobra"

var (
	showCmd = &cobra.Command{
		Use:   "show",
		Short: "Show last N rows in results",
		Run:   runShowCommand,
	}

	euroJackpotCmd = &cobra.Command{
		Use:   "eurojackpot",
		Short: "Eurojackpot set of commands",
	}
)

func RegisterNewCommand(rootCmd *cobra.Command) {
	euroJackpotCmd.AddCommand(showCmd)

	rootCmd.AddCommand(euroJackpotCmd)
}
