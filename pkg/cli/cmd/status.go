package cmd

import (
	"github.com/rutmanz/hackhour-go/pkg/api"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the HackHour Api Status",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := api.NewHackHourClient("")
		status, err := client.Status()
		if err != nil {
			return err
		}
		getJsonEncoder().Encode(status)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
