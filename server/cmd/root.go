package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Todo server 1",
	Long:  "All software has version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Todo API")
	},
}

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "root command",
	Long:  `we'll see !`,
	Run: func(cmd *cobra.Command, args []string) {
		// ...
	},
}

func Execute() {
	if err := versionCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
