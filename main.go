package main

import (
	"github.com/babanini95/diaper-duty/cmd"
	_ "modernc.org/sqlite"
)

func main() {
	cmd.Execute()
}
