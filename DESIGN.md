# Diaper Duty

CLI tool for tracking diaper changes for your baby.

## Objective

To create a standalone command-line application that allows a user to initialize a baby profile, log diaper changes, and check the status and history. All data will be stored locally in a single SQL database file (like SQLite), making it simple and portable.

## Schema

1. Table `profile`

    This table will have only a single row to store the baby's information and custom settings.

    | Column Name             | Type           | Description                                                                               |
    | ----------------------- | -------------- | ----------------------------------------------------------------------------------------- |
    | id                      | INTEGER        | Primary key (only 1)                                                                      |
    | baby_name               | TEXT           | Baby's name                                                                               |
    | baby_birthdate          | TEXT           | Baby's birthdate in YYYY-MM-DD format                                                     |
    | diaper_interval_minutes | INTEGER (NULL) | Custom interval for diaper change reminders in minutes. If `NULL`, use age-based interval |

2. Table `changes`

    This table will store all diaper change records.

    | Column Name | Type    | Description                                                      |
    | ----------- | ------- | ---------------------------------------------------------------- |
    | id          | INTEGER | Primary key                                                      |
    | change_time | TEXT    | Timestamp of the change in ISO 8601 format (YYYY-MM-DDTHH:MM:SS) |
    | notes       | TEXT    | Optional notes about the change                                  |

## Commands

### `diaper-duty init`

* **Purpose:** The one-time setup command.
* **Action:**
  * Checks if the database/profile already exists. If so, it exits with an error.
  * If not, it launches an interactive prompt for the baby's name and birthday.
  * Creates the database, tables, and saves the initial profile info.
* **Example:**
  
  ```bash
  $ diaper-duty init
  Welcome! Let's set up your Diaper Duty Tracker.
  > Baby's name: Alice
  > Baby's birthday (YYYY-MM-DD): 2023-01-01
  Success! Profile for Alice created. You're ready to 'diaper-duty log' the first change!
  ```

### `diaper-duty log`

* **Purpose:** To record a new diaper change.
* **Action:** Inserts a new row into the changes table with the current timestamp.
* **Flags:**
  * `--note (some text)` or `-n (some text)` (optional): To add an optional note to the log entry.
* **Example:**

  ```bash
  $ diaper-duty log --note "A big one"
  ✅ Logged diaper change for Alice at 00:31 AM. Note: A big one.
  ```

### `diaper-duty status`

* **Purpose:** To check the current diaper change status.
* **Actions:**
  * Gets most recent change from `changes` table.
  * Gets the baby's profile from `profile` table.
  * Calculates the time since the last change.
  * Calculates the next recommended change time based on the baby's age or custom interval.
  * Displays the status in a user-friendly format.
* **Example:**
  
  ```bash
  $ diaper-duty status
  --- Diaper Status for Alice ---
  Last change: 45 minutes ago (at 00:31 AM)
  Reminder cycle: 2 hours (based on age)
  Next change due: in 1 hour 15 minutes (at 02:31 AM)
  ```

### `diaper-duty history`

* **Purpose:** To view today history of diaper changes
* **Actions:** Queries the `changes` table for all entries from the current day, ordered from most to least recent.
* **Example:**

  ```bash
  $ diaper-duty history
  --- Alice history for today (Tuesday, 15 Jul 2025) ---
  - 11:46 PM (45m ago) - Note: Big one!
  - 09:30 PM
  - 07:15 PM
  - 05:00 PM
  ```

### `diaper-duty config`

* **Purpose:** To allow users customize the reminder interval.
* **Actions:** Updates the `reminder_interval_minutes` in the profile table.
* **Flags:**
  * `--set-reminder <duration>`: Sets a custom interval (e.g., 2h, 90m, 2h30m)
  * `--reset-reminder`: Resets the interval back to NULL, re-enabling age-based defaults.
* **Example:**

  ```bash
  $ diaper-duty config --set-reminder 2h30m
  Success! Reminder interval set to 2 hours and 30 minutes.

  $ diaper-duty config --reset-reminder
  Success! Reminder reset to age-based defaults.
  ```
