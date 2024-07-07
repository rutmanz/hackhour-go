/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := newClient()
		historyPtr, err := client.GetHistory()
		if err != nil {
			fmt.Printf("Failed to get history: %v\n", err)
			return
		}
		history := *historyPtr
		sort.Slice(history, func(i, j int) bool {
			if cmd.Flag("reverse").Changed {
				return history[i].CreatedAt.After(history[j].CreatedAt)
			}
			return history[i].CreatedAt.Before(history[j].CreatedAt)
		})
		err = getJsonEncoder().Encode(history)
		if err != nil {
			fmt.Println(err, history)
			return
		}
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
