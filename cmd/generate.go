package cmd

import (
	"fmt"
	"os"

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
		generateLicenseFile()
	},
}

func generateLicenseFile() {
	f, err := os.Create("LICENSE")

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	defer f.Close()

	_, err2 := f.WriteString("Hello, world!")

	if err2 != nil {
		fmt.Printf("%v", err)
		return
	}
}