package cmd

import (
	"file-organizer-cli/organizer"
	"fmt"

	"github.com/spf13/cobra"
)

var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Preview how files would be organized without moving them",
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
		err = organizer.PreviewFiles(targetDir, rules)
		if err != nil {
			fmt.Println("Error previewing files:", err)
		}
	},
}
