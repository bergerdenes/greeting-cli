package cmd

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all persons",
	Run: func(cmd *cobra.Command, args []string) {
		client := resty.New()
		resp, err := client.R().Get("http://localhost:8080/api/people")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Response:", resp.String())
	},
}
