package cmd

import (
	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "A note for cobra",
	Long:  `A long cobra note`,
}

func init() {
	rootCmd.AddCommand(noteCmd)
}
