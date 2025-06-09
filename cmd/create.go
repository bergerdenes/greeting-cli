package cmd

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var name string
var age int
var gender string
var keywords []string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new person",
	Run: func(cmd *cobra.Command, args []string) {
		client := resty.New()

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(map[string]interface{}{
				"name":     name,
				"age":      age,
				"gender":   gender,
				"keywords": keywords,
			}).
			Post("http://localhost:8080/api/people")

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Response:", resp.String())
	},
}

func init() {
	createCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the person")
	createCmd.Flags().IntVarP(&age, "age", "a", 0, "Age of the person")
	createCmd.Flags().StringVarP(&gender, "gender", "g", "", "Gender (MALE, FEMALE, OTHER)")
	createCmd.Flags().StringSliceVarP(&keywords, "keywords", "k", []string{}, "Keywords")

	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("age")
	createCmd.MarkFlagRequired("gender")
}
