/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancels the current HackHour session",
	Run: func(cmd *cobra.Command, args []string) {
		client := newClient()
		session, err := client.SessionCancel()
		if err != nil {
			fmt.Println("Failed to cancel session:", err)
			return
		}
		fmt.Println("Session cancelled:", session.ID)
		getJsonEncoder().Encode(session)
	},
}

func init() {
	sessionCmd.AddCommand(cancelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cancelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cancelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
