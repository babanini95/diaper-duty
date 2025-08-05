package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the current diaper change status",
	Long:  `This command will give you info about the last time your baby change diaper, your reminder cycle, and the next change due`,
	RunE:  commandHandler,
}

func commandHandler(cmd *cobra.Command, args []string) error {
	// Fetch profileData data
	profileData, err := appState.db.GetProfile(cmd.Context())
	if err != nil {
		return fmt.Errorf("can not get profile: %v", err)
	}

	// Fetch the last change data
	lastChangeData, err := appState.db.GetTheLastChange(cmd.Context())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("you haven't log the diaper change")
		}
		return fmt.Errorf("can not get the last change data: %v", err)
	}

	// Parse last change
	lastChange, _ := time.Parse(time.RFC3339, lastChangeData.ChangeTime)

	// Parse baby birthday
	babyBirthday, _ := time.Parse(time.DateOnly, profileData.BabyBirthday)

	// Calculate between now and the last change
	timeSince := time.Since(lastChange)

	// Calculate time remaining
	var changeIntervalInMinutes int64
	intervalNote := ""
	if profileData.DiaperIntervalMinutes.Valid {
		changeIntervalInMinutes = profileData.DiaperIntervalMinutes.Int64
		intervalNote = "(custom)"
	} else {
		// Calculate baby's age in month
		babyAge := calculateAgeInMonth(babyBirthday)

		if babyAge <= 2 {
			changeIntervalInMinutes = 2 * 60
		} else if babyAge <= 6 {
			changeIntervalInMinutes = 3 * 60
		} else {
			changeIntervalInMinutes = 4 * 60
		}
		intervalNote = "(default based on age)"
	}

	timeRemaining := changeIntervalInMinutes*int64(time.Minute) - int64(timeSince)

	// Construct result
	lastChangeText := formatMinute(timeSince.Round(time.Minute).Minutes())
	nextChangeDue := time.Duration(timeRemaining).Round(time.Minute).Minutes()
	nextChangeDueText := ""
	if nextChangeDue < 0 {
		nextChangeDueText = fmt.Sprintf("It's %v ago, you're late!!", formatMinute(-1*nextChangeDue))
	} else {
		nextChangeDueText = fmt.Sprintf("In %v", formatMinute(nextChangeDue))
	}

	// Print result
	fmt.Printf("--- Diaper status for %s ---\n", profileData.BabyName)
	fmt.Printf(`Last Change:     %v ago (at %v)
Reminder Cycle:  %v hours %v
Next Change Due: %s
`,
		lastChangeText,
		lastChange.Local().Format(time.Kitchen),
		changeIntervalInMinutes/60,
		intervalNote,
		nextChangeDueText,
	)

	return nil
}

func calculateAgeInMonth(birthdate time.Time) int {
	now := time.Now()
	years := now.Year() - birthdate.Year()
	months := now.Month() - birthdate.Month()

	if now.Day() < birthdate.Day() {
		months--
	}

	ageInMonth := years*12 + int(months)
	return ageInMonth
}

func formatMinute(minute float64) string {
	hours := math.Trunc(minute / 60)
	minutes := math.Mod(minute, 60)
	if hours < 0 {
		minutes *= -1
	}

	return fmt.Sprintf("%v hour(s) %v minute(s)", hours, minutes)
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
