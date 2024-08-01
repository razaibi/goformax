package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of goformax",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("goformax version 0.0.1")
	},
}
