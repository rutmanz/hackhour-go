/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

// goalsCmd represents the goals command
var goalsCmd = &cobra.Command{
	Use:   "goals",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := newClient()
		goalsPtr, err := client.GetGoals()
		if err != nil {
			fmt.Printf("Failed to get history: %v\n", err)
			return
		}
		goals := *goalsPtr
		sort.Slice(goals, func(i, j int) bool {
			if cmd.Flag("desc").Changed {
				return goals[i].Minutes > goals[j].Minutes
			}
			return goals[i].Minutes < goals[j].Minutes
		})
		err = getJsonEncoder().Encode(goals)
		if err != nil {
			fmt.Println(err, goals)
			return
		}
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
