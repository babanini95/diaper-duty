package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/babanini95/diaper-duty/db/migration"
	"github.com/babanini95/diaper-duty/internal/database"
	"github.com/pressly/goose/v3"
)

type state struct {
	db *database.Queries
}

var appState *state

func CreateQueries() error {
	// Create database filepath
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not get user home dir: %v", err)
	}
	dbPath := filepath.Join(home, ".diaper-duty", "diaper-duty.db")

	// Ensure the directory exists
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return fmt.Errorf("could not create db directory: %v", err)
	}

	fmt.Println("database path: " + dbPath)

	db, err := sql.Open("sqlite", "file:"+dbPath)
	if err != nil {
		return fmt.Errorf("queries can not be created: %v", err)
	}

	// Upgrade database to the most updated migration
	gooseProvider, err := goose.NewProvider(goose.DialectSQLite3, db, migration.Embed)
	if err != nil {
		return err
	}

	_, err = gooseProvider.Up(context.Background())
	if err != nil {
		return err
	}

	appState.db = database.New(db)
	return nil
}

func CreateNewState() *state {
	return &state{}
}
