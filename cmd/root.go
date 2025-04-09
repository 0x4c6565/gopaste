package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const APIURL = "https://p.lee.io/api/v1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gopaste",
	Short: "CLI client for p.lee.io paste service",
	Long:  `A CLI tool for interacting with the p.lee.io paste service with client-side encryption support`,
}

func Execute() {
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	rootCmd.AddCommand(createPasteCommand())
	rootCmd.AddCommand(getPasteCommand())
	rootCmd.AddCommand(listSyntaxCommand())
	rootCmd.AddCommand(listExpiresCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
