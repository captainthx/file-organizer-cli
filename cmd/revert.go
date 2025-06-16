package cmd

import (
	"file-organizer-cli/organizer"
	"fmt"

	"github.com/spf13/cobra"
)

var revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Revert the last organization by restoring file paths",
	Run: func(cmd *cobra.Command, args []string) {
		err := organizer.RevertLastMove()
		if err != nil {
			fmt.Println("Error reverting moves: ", err)
		}
	},
}
