/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"sort"

	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:     "history",
	GroupID: "data",
	Short:   "Shows your HackHour session history",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := newClient()
		historyPtr, err := client.GetHistory()
		if err != nil {
			return err
		}
		history := *historyPtr
		sort.Slice(history, func(i, j int) bool {
			if cmd.Flag("reverse").Changed {
				return history[i].CreatedAt.After(history[j].CreatedAt)
			}
			return history[i].CreatedAt.Before(history[j].CreatedAt)
		})
		getJsonEncoder().Encode(history)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// historyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	historyCmd.Flags().BoolP("reverse", "r", false, "Sort with newest at the top")
}
