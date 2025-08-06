# Diaper Duty Tracker

A simple cli tool to track diaper changes for your baby.

## Motivation

When I am entering the programming mode, I want my concentration 100% distributed to either my browser, my editor, or my terminal. But as a father of a cute (but poop a lot) little baby, my wife want me to involve in the parenting especially to remember when the last diaper change was done, and when the next one is due. So I created this tool to help me track diaper changes without leaving my coding environment.

## Tech Stack

* **Language:** Go
* **CLI Framework:** [Cobra](https://github.com/spf13/cobra)
* **Database:** SQLite
* **Migrations:** [Goose](https://github.com/pressly/goose)
* **Queries:** [sqlc](https://github.com/sqlc-dev/sqlc) for type safe SQL code generation.

## Features

* **Interactive Setup:** A one time `init` command to set up your baby's profile.
* **Quick Logging:** Easily `log` a new diaper change, with options to specify the time and add notes.
* **Instant Status:** Get an immediate `status` update showing the last change and when the next one is due.
* **Custom Configuration:** Use the `config` command to set your own custom reminder intervals.
* **Daily History:** View a clean `history` of all changes logged for the current day.

## Getting Started

### Prerequisites

* Go (version 1.21 or later)
* [Goose CLI](https://github.com/pressly/goose?tab=readme-ov-file#install) installed (`go install github.com/pressly/goose/v3/cmd/goose@latest`)

### Installation & Setup

1. **Clone the repository:**

    ```bash
    git clone [https://github.com/your-username/diaper-duty.git](https://github.com/your-username/diaper-duty.git)
    cd diaper-duty
    ```

2. **Install dependencies:**

    ```bash
    go mod tidy
    ```

3. **Copy the example environment file:**

    ```bash
    cp .env.example .env
    ```

    Update the `.env` file with your desired `goose` [environment settings](https://pressly.github.io/goose/documentation/environment-variables/).

4. **Run the database migrations:**
    This command creates the `diaper-duty.db` file and sets up the necessary tables.

    ```bash
    goose up
    ```

5. **Build the application:**

    ```bash
    go build -o diaper-duty .
    ```

6. **Initialize the tracker:**
    Run the interactive setup to create your baby's profile.

    ```bash
    ./diaper-duty init
    ```

---

## Usage Examples

Once initialized, you can use the following commands:

* **Log a change for the current time:**

    ```bash
    ./diaper-duty log -n "A big one"
    ```

* **Log a change for a specific time in the past:**

    ```bash
    ./diaper-duty log -t 2:30PM
    ```

* **Check the current status:**

    ```bash
    ./diaper-duty status
    ```

* **View today's history:**

    ```bash
    ./diaper-duty history
    ```

* **Set a custom reminder interval:**

    ```bash
    ./diaper-duty config --set-reminder 2h30m
    ```
