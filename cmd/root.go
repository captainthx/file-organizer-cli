/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "organizer",
	Short: "A CLI tool to organize your files by extension",
}

func Execute() {
	rootCmd.AddCommand(organizeCmd)
	rootCmd.AddCommand(previewCmd)
	rootCmd.AddCommand(revertCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}

var targetDir string
var rulesFile string

func init() {
	rootCmd.PersistentFlags().StringVarP(&targetDir, "dir", "d", "", "Target directory")
	rootCmd.PersistentFlags().StringVarP(&rulesFile, "rules", "r", "rules.json", "Rules file path")

	organizeCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Confirm before moving files")
}
