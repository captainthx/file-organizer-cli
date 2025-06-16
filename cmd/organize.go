package cmd

import (
	"file-organizer-cli/organizer"
	"fmt"

	"github.com/spf13/cobra"
)

var interactive bool
var organizeCmd = &cobra.Command{
	Use:   "run",
	Short: "Organize files by extension",
	Run: func(cmd *cobra.Command, args []string) {
		if targetDir == "" {
			fmt.Println("Please specify a target directory using --dir")
			return
		}
		rules, err := organizer.LoadRules(rulesFile)
		if err != nil {
			fmt.Println("Error loading rules: ", err)
			return
		}

		err = organizer.OrganizeFiles(targetDir, rules, interactive)
		if err != nil {
			fmt.Println("Error organizing file: ", err)
		}

	},
}
