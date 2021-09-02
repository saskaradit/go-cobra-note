package main

import (
	"github.com/saskaradit/go-cobra/cmd"
	"github.com/saskaradit/go-cobra/internal/db"
)

func main() {
	db.OpenDatabase()
	cmd.Execute()
}
