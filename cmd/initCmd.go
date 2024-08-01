package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize project.",
	Run: func(cmd *cobra.Command, args []string) {

		fileName := "data.json"
		content := `{
    "backend" : {
      "model" : "Sample"
    },
    "form": {
      "header": "Sample Form",
      "tagline" : "A simple form to enter data",
      "action": "http://localhost:3000/user",
      "method": "POST",
      "elements": [
        {
          "type": "text",
          "name": "name",
          "label": "Name",
          "labelStyle" : "text-primary",
          "hint": "Enter your name without special characters.",
          "placeholder": "Name",
          "required": true
        },
        {
          "type": "number",
          "name": "age",
          "label": "Age",
          "labelStyle" : "text-primary",
          "hint": "Enter your age.",
          "placeholder": "Age",
          "required": true
        },
        {
          "type": "email",
          "name": "email",
          "label": "Email",
          "labelStyle" : "text-primary",
          "hint": "Enter your email.",
          "placeholder": "Email",
          "required": true
        },
        {
          "type": "checkbox",
          "name": "rememberMe",
          "label": "Remember Me"
        },
        {
          "type": "checkbox-tile",
          "name": "sampleTile",
          "label": "Checkbox Tiles",
          "checks": [
            { "value": "check1", "label": "Check 1" },
            { "value": "check2", "label": "Check 2" }
          ]
        },
        {
          "type": "radio",
          "name": "gender",
          "label": "Gender",
          "options": [
            { "value": "male", "label": "Male" },
            { "value": "female", "label": "Female" }
          ]
        },
        {
          "type": "radio-tile",
          "name": "choices",
          "label": "Choices",
          "options": [
            { "value": "opt1", "label": "1" },
            { "value": "opt2", "label": "2" }
          ]
        },
        {
          "type": "list-tile",
          "name": "listTiles",
          "label": "List Tiles",
          "labelStyle" : "text-primary mt-2",
          "items": [
            { "value": "opt1", "label": "Item 1" },
            { "value": "opt2", "label": "Item 2" }
          ]
        },
        {
          "type": "poll",
          "name": "poll",
          "label": "Items Poll",
          "labelStyle" : "mt-2",
          "totalCount": "7",
          "items": [
            { "value": "opt1", "label": "Item 1", "count": "6" },
            { "value": "opt2", "label": "Item 2", "count": "1"  }
          ]
        }
      ]
    }
  }`

		// Call the function to create the file with content
		err := createFileWithContent(fileName, content)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

	},
}

func createFileWithContent(fileName string, content string) error {
	// Create or open the file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
