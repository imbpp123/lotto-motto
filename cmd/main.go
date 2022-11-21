package main

import (
	"fmt"
	"github.com/imbpp123/lotto_motto/cmd/eurojackpot"
	"github.com/imbpp123/lotto_motto/cmd/lotto6aus49"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "lotto",
	Short: "Commands that use https://www.lotto-berlin.de/ data",
}

func main() {
	lotto6aus49.RegisterNewCommand(rootCmd)
	eurojackpot.RegisterNewCommand(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
