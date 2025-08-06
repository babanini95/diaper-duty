package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "View today history of diaper changes",
	Long:  `Display all of diaper changes log for today`,
	Run:   historyCommandHandler,
}

func historyCommandHandler(cmd *cobra.Command, args []string) {
	// Get today history data
	now := time.Now()
	changes, err := appState.db.ListHistoryByDate(cmd.Context(), now.Format(time.RFC3339))
	if err != nil {
		log.Fatalf("Failed to get history data: %v", err)
	}

	if len(changes) == 0 {
		log.Fatal("No changes have been logged today")
	}

	/*
		--- History for Today (Wednesday, August 6) ---
		- 9:15 AM (1h 32m ago)
		- 7:30 AM (3h 17m ago) - Note: Woke up
		- 4:00 AM (6h 47m ago)
	*/

	// Build output
	fmt.Printf("--- History for Today (%v) ---\n", now.Format("Monday, January 2"))
	for _, change := range changes {
		t, _ := time.Parse(time.RFC3339, change.ChangeTime)
		timeSince := formatMinute(time.Since(t).Round(time.Minute).Minutes())
		note := ""
		if change.Notes.Valid {
			note = fmt.Sprintf("- Note: %s", change.Notes.String)
		}

		fmt.Printf("- %s (%s ago) %s\n", t.Local().Format(time.Kitchen), timeSince, note)
	}
}

func init() {
	rootCmd.AddCommand(historyCmd)
}
