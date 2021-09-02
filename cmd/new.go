package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/saskaradit/go-cobra/internal/db"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new note",
	Long:  `Creates a new studybuddy note`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewNote()
	},
}

type promptContent struct {
	errorMsg string
	label    string
}

func init() {
	noteCmd.AddCommand(newCmd)
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	res, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt Failed", err)
	}

	fmt.Println("Input:", res)
	return res
}

func promptGetSelect(pc promptContent) string {
	items := []string{"animal", "food", "person", "object"}
	index := -1

	var res string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}
		index, res, err = prompt.Run()
		if index == -1 {
			items = append(items, res)
		}
	}

	if err != nil {
		fmt.Println("Prompt Failed ", err)
		os.Exit(1)
	}
	fmt.Println("input: ", res)
	return res
}

func createNewNote() {
	wordPromptContent := promptContent{
		"Please provide a word",
		"What word would you like to make a note of? ",
	}
	word := promptGetInput(wordPromptContent)
	definitionPromptContent := promptContent{
		"Please provide a definition",
		fmt.Sprintf("What is the definition of %s? ", word),
	}
	definition := promptGetInput(definitionPromptContent)

	categoryPromptContent := promptContent{
		"Please provide a category",
		fmt.Sprintf("What category does %s belongs to? ", word),
	}
	category := promptGetSelect(categoryPromptContent)

	db.InsertNote(word, definition, category)
}
