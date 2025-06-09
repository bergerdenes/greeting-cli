package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "greeting-cli",
	Short: "CLI tool to manage Person entities via OpenAPI3 endpoint",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(updateCmd)
}
