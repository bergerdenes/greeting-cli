package cmd

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var deleteID string

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a person by ID",
	Run: func(cmd *cobra.Command, args []string) {
		client := resty.New()
		url := fmt.Sprintf("http://localhost:8080/api/people/%s", deleteID)
		resp, err := client.R().Delete(url)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if resp.IsError() {
			fmt.Println("Error:", resp.RawResponse.Status)
			return
		}
		fmt.Println("Response: Deleted")
	},
}

func init() {
	deleteCmd.Flags().StringVarP(&deleteID, "id", "i", "", "ID of the person to delete")
	deleteCmd.MarkFlagRequired("id")
}
