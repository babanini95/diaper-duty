package cmd

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Customize the reminder interval",
	Long: `Use this if you want to change the reminder interval.
By default, the reminder interval will goes by your baby's age.
- 0 - 2 months old will set 2 hours interval
- >2 - 6 months old will set 3 hours interval
- >6 months old will set 4 hours interval`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the flags value
		setReminderFlag, _ := cmd.Flags().GetString("set-reminder")
		resetReminderFlag, _ := cmd.Flags().GetBool("reset-reminder")

		if setReminderFlag == "" && !resetReminderFlag {
			profile, err := appState.db.GetProfile(cmd.Context())
			if err != nil {
				return fmt.Errorf("failed to get your config. %s", err.Error())
			}
			if !profile.DiaperIntervalMinutes.Valid {
				fmt.Println("You used the default interval.",
					"\n- 0 - 2 months old will set 2 hours interval",
					"\n- >2 - 6 months old will set 3 hours interval",
					"\n- >6 months old will set 4 hours interval",
					"\nUse -s, --set-reminder <interval> to set a new interval",
				)
				return nil
			}

			fmt.Println("Your will be reminded each",
				formatMinute(float64(profile.DiaperIntervalMinutes.Int64)),
				"\nUse: \n-s, --set-reminder <interval> to set a new interval or \n-r, --reset-reminder to use default interval",
			)
			return nil
		}

		// Set new interval if needed
		if setReminderFlag != "" {
			// Parse new interval
			newInterval, err := time.ParseDuration(setReminderFlag)
			if err != nil {
				return fmt.Errorf("wrong interval format. %s", err.Error())
			}

			// Save in minute format
			err = appState.db.SetCustomReminder(cmd.Context(), sql.NullInt64{
				Valid: true,
				Int64: int64(newInterval.Minutes()),
			})
			if err != nil {
				return fmt.Errorf("failed to save the config. %s", err.Error())
			}
		}

		// Reset interval if needed
		if resetReminderFlag {
			err := appState.db.ResetCustomReminder(cmd.Context())
			if err != nil {
				return fmt.Errorf("failed to reset config. %s", err.Error())
			}
		}

		fmt.Println("✅ Successfully saved your config")
		return nil
	},
}

func init() {
	configCmd.Flags().StringP("set-reminder", "s", "", "Set the reminder interval. Format: 1h, 90m, 1h30m")
	configCmd.Flags().BoolP("reset-reminder", "r", false, "Reset the reminder interval to use the default reminder")
	rootCmd.AddCommand(configCmd)
}
