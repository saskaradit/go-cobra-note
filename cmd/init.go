package cmd

import (
	"github.com/saskaradit/go-cobra/internal/db"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Go Cobra from init",
	Long:  `Lorem From Init`,
	Run: func(cmd *cobra.Command, args []string) {
		db.CreateTable()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
