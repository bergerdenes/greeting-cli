package cmd

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var greetID string

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greet a person by ID",
	Run: func(cmd *cobra.Command, args []string) {
		client := resty.New()
		url := fmt.Sprintf("http://localhost:8080/api/people/%s/greeting", greetID)
		resp, err := client.R().Get(url)
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
	greetCmd.Flags().StringVarP(&greetID, "id", "i", "", "ID of the person to greet")
	greetCmd.MarkFlagRequired("id")
}
