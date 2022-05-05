package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new license file",
	Long:  `Generate a new license file based on your own need`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("License file generated")
	},
}
