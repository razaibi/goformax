package cmd

import (
	"encoding/json"
	"fmt"
	"goformax/logic"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generates an HTML form with Spectre CSS.",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.ReadFile("data.json")
		if err != nil {
			fmt.Print(err)
		}

		// Unmarshall data.
		var data logic.JSONData
		err = json.Unmarshal([]byte(file), &data)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}

		content, _ := logic.GenerateHTML(data)

		err = os.WriteFile("index.html", []byte(content), 0644)
		if err != nil {
			log.Fatal("Error while generating file. ", err)
		}
	},
}
