/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"sort"

	"github.com/spf13/cobra"
)

// goalsCmd represents the goals command
var goalsCmd = &cobra.Command{
	Use:     "goals",
	GroupID: "data",
	Short:   "Shows your HackHour goals list",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := newClient()
		goalsPtr, err := client.GetGoals()
		if err != nil {
			return err
		}
		goals := *goalsPtr
		sort.Slice(goals, func(i, j int) bool {
			if cmd.Flag("desc").Changed {
				return goals[i].Minutes > goals[j].Minutes
			}
			return goals[i].Minutes < goals[j].Minutes
		})
		getJsonEncoder().Encode(goals)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(goalsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goalsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	goalsCmd.Flags().BoolP("desc", "d", false, "Sorts goals in descending order")
}
