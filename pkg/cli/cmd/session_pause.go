/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pauseCmd represents the pause command
var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pauses or resumes the current session",
	Run: func(cmd *cobra.Command, args []string) {
		client := newClient()
		session, err := client.SessionPause()
		if err != nil {
			fmt.Println("Failed to pause session:", err)
			return
		}
		if session.Paused {
			fmt.Println("Session paused:", session.ID)
		} else {
			fmt.Println("Session resumed:", session.ID)
		}

		getJsonEncoder().Encode(session)
	},
}

func init() {
	sessionCmd.AddCommand(pauseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pauseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pauseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
