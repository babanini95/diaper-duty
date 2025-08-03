package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "diaper-duty",
	Short: "A CLI tool to track baby diaper changes.",
	Long: `Diaper Duty is a simple CLI tool to help you remember
when your baby's diaper was last changed.`,
}

func Execute() {
	appState = CreateNewState()
	err := CreateQueries()
	if err != nil {
		fmt.Printf("Cannot set database: %v", err)
		os.Exit(1)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
