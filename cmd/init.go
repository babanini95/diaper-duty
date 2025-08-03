package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/babanini95/diaper-duty/internal/database"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the Diaper Duty Tracker with your baby's info.",
	Long:  `This command sets up the tracker for the first time. It will create the necessary database file and prompt you for your baby's name and birthday.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing Diaper Duty Tracker...")

		// Check if the profile already exists. If so, exit.
		c, err := appState.db.CountProfiles(rootCmd.Context())
		if err != nil {
			log.Fatalf("Can not check database: %v", err)
		}
		if c > 0 {
			fmt.Println("Already initialized")
			return
		}

		// Prompt for baby's name and birthday.
		promptName := promptui.Prompt{
			Label: "Baby's name",
			Validate: func(s string) error {
				if s == "" {
					return fmt.Errorf("name must not be empty")
				}
				return nil
			},
		}

		promptBirthday := promptui.Prompt{
			Label: "Baby's birthday (YYYY-MM-DD)",
			Validate: func(s string) error {
				_, err := time.Parse(time.DateOnly, s)
				if err != nil {
					return fmt.Errorf("wrong birthday format")
				}
				return nil
			},
		}

		babysName, err := promptName.Run()
		if err != nil {
			log.Fatalf("Can not create name: %v", err)
		}
		babysBirthday, err := promptBirthday.Run()
		if err != nil {
			log.Fatalf("Can not add birthday: %v", err)
		}

		// Save the info.
		profile, err := appState.db.CreateProfile(rootCmd.Context(), database.CreateProfileParams{
			BabyName:     babysName,
			BabyBirthday: babysBirthday,
		})
		if err != nil {
			log.Fatalf("Failed to save baby's profile: %v", err)
		}

		// Print a success message.
		fmt.Printf("Successfully created %s profile, Happy Duty!!\n", profile.BabyName)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
