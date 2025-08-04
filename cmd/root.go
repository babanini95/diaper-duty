package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "diaper-duty",
	Short: "A CLI tool to track baby diaper changes.",
	Long: `Diaper Duty is a simple CLI tool to help you remember
when your baby's diaper was last changed.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Name() == "help" || cmd.Name() == "completion" || cmd.Name() == "init" {
			return nil
		}
		count, err := appState.db.CountProfiles(context.Background())
		if err != nil {
			return fmt.Errorf("database check failed: %w", err)
		}

		if count == 0 {
			return fmt.Errorf("project not initialized. Please run 'diaper-duty init' first")
		}
		return nil
	},
}

func Execute() {
	appState = CreateNewState()
	err := CreateQueries()
	if err != nil {
		log.Fatalf("Cannot set database: %v", err)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
