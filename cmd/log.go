package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/babanini95/diaper-duty/internal/database"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Log a new diaper change",
	Long:  `Records a new diaper change. Defaults to the current time unless a specific time is provided with the --time flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get flags string
		noteFlag, _ := cmd.Flags().GetString("note")
		timeFlag, _ := cmd.Flags().GetString("time")

		// Parse time flag
		var changeTime time.Time
		now := time.Now()

		if timeFlag != "" {
			layout := "3:15PM"
			parsedTime, err := time.Parse(layout, timeFlag)
			if err != nil {
				log.Fatalln("Wrong time format")
			}

			changeTime = time.Date(now.Year(), now.Month(), now.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, now.Location())
		} else {
			changeTime = now
		}

		// Save to database
		param := database.InsertDiaperChangeParams{
			Notes:      sql.NullString{String: noteFlag, Valid: noteFlag != ""},
			ChangeTime: changeTime.Format(time.RFC3339),
		}
		_, err := appState.db.InsertDiaperChange(rootCmd.Context(), param)
		if err != nil {
			log.Fatalf("Failed to log the change: %v", err)
		}
		fmt.Println("✅ Logged diaper change at", changeTime.Format("3:04 PM"))
	},
}

func init() {
	logCmd.Flags().StringP("note", "n", "", "Add an optional note to the log.")
	logCmd.Flags().StringP("time", "t", "", "Log a change at a specific time (e.g., '3:15PM')")
	rootCmd.AddCommand(logCmd)
}
