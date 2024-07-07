/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:     "stats",
	GroupID: "data",
	Short:   "Shows your HackHour statistics",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := newClient()
		stats, err := client.GetStats()
		if err != nil {
			fmt.Printf("Failed to get stats: %v\n", err)
			return
		}
		err = getJsonEncoder().Encode(stats)
		if err != nil {
			fmt.Println(err, stats)
			return
		}
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
