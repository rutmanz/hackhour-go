package cmd

import (
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:     "stats",
	GroupID: "data",
	Short:   "Shows your HackHour statistics",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := newClient()
		stats, err := client.GetStats()
		if err != nil {
			return err
		}
		getJsonEncoder().Encode(stats)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
