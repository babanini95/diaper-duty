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

	gooseProvider, err := goose.NewProvider(goose.DialectSQLite3, db, migration.Embed)
	if err != nil {
		return err
	}

	results, err := gooseProvider.Up(context.Background())
	if err != nil {
		return err
	}

	for _, r := range results {
		fmt.Printf("%v %v done: %v\n", r.Source.Type, r.Source.Version, r.Duration)
	}

	appState.db = database.New(db)
	// appState.db.CreateProfilesTable(context.Background())
	// appState.db.CreateChangesTable(context.Background())
	return nil
}

func CreateNewState() *state {
	return &state{}
}
