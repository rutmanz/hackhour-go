/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statasCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statasCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
