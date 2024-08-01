package main

import (
	"fmt"
	"os"

	"goformax/cmd"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "mycli"}

	// Add commands to rootCmd
	rootCmd.AddCommand(cmd.InitCmd)
	rootCmd.AddCommand(cmd.HelloCmd)
	rootCmd.AddCommand(cmd.VersionCmd)
	rootCmd.AddCommand(cmd.GenerateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
