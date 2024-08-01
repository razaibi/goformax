package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var HelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints a greeting message",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, Welcome to goformax!")
	},
}
