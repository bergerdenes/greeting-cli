package cmd

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var updateID string

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a person by ID",
	Run: func(cmd *cobra.Command, args []string) {
		client := resty.New()
		url := fmt.Sprintf("http://localhost:8080/api/people")

		body := make(map[string]interface{})

		body["id"] = updateID

		if cmd.Flags().Changed("name") {
			body["name"] = name
		}
		if cmd.Flags().Changed("age") {
			body["age"] = age
		}
		if cmd.Flags().Changed("gender") {
			body["gender"] = gender
		}
		if cmd.Flags().Changed("keywords") {
			body["keywords"] = keywords
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Patch(url)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if resp.IsError() {
			fmt.Println("Error:", resp.RawResponse.Status)
			return
		}
		
		fmt.Println("Response:", resp.String())
	},
}

func init() {
	updateCmd.Flags().StringVarP(&updateID, "id", "i", "", "ID of the person to update")
	updateCmd.Flags().StringVarP(&name, "name", "n", "", "New name")
	updateCmd.Flags().IntVarP(&age, "age", "a", 0, "New age")
	updateCmd.Flags().StringVarP(&gender, "gender", "g", "", "New gender")
	updateCmd.Flags().StringSliceVarP(&keywords, "keywords", "k", []string{}, "New keywords")
	updateCmd.MarkFlagRequired("id")
}
