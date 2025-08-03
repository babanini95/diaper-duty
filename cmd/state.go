package cmd

import (
	"database/sql"
	"fmt"

	"github.com/babanini95/diaper-duty/internal/database"
)

type state struct {
	db *database.Queries
}

var appState *state

func CreateQueries() error {
	db, err := sql.Open("sqlite", "diaper-duty.db")
	if err != nil {
		return fmt.Errorf("queries can not be created: %v", err)
	}

	appState.db = database.New(db)
	return nil
}

func CreateNewState() *state {
	return &state{}
}
