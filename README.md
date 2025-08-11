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

The initial version of the application is a complete, standalone CLI tool with a full set of features:

* **Interactive Setup:** A one time `init` command to create a profile for your baby.
* **Smart Logging:** Quickly `log` a diaper change, with flags to specify a past time (`-t "2:30PM"`) or add notes (`-n "..."`).
* **Instant Status:** The `status` command provides an immediate summary of the last change and calculates when the next one is due based on the baby's age.
* **Daily History:** The `history` command gives a clean overview of all changes logged for the day.
* **Easy Configuration:** Users can override the smart defaults and `config` a custom reminder interval.

## Getting Started

### Prerequisites

* Go (version 1.21 or later)
* [Goose CLI](https://github.com/pressly/goose?tab=readme-ov-file#install) installed (`go install github.com/pressly/goose/v3/cmd/goose@latest`)

### Installation & Setup

Install using Go:

```bash
go install github.com/babanini95/diaper-duty@latest
```

or build from source:

1. **Clone the repository:**

    ```bash
    git clone https://github.com/babanini95/diaper-duty.git
    cd diaper-duty
    ```

2. **Install dependencies:**

    ```bash
    go mod tidy
    ```

3. **Build the application:**

    ```bash
    go build -o diaper-duty .
    ```

4. **Initialize the tracker:**
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

## Future Plans

The next major step is to evolve the application's architecture by building a background **API server** and a simple, mobile friendly **Web UI**. This will allow non technical users (like my wife!) to use the tracker from their phone's web browser, while the CLI remains fully functional for developer users like me.
