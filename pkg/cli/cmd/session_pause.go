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
	Short: "Pauses the current session",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := newClient()
		session, err := client.GetSession()
		if err != nil {
			return err
		}
		if session.Paused {
			fmt.Println("Session already paused")
		} else {
			fmt.Println("Pausing session...")
			res, err := client.SessionPause()
			if err != nil {
				return err
			}
			if res.Paused {
				fmt.Println("Session paused")
			} else {
				fmt.Println("Session could not be paused")
			}
		}
		return nil
	},
}

func init() {
	sessionCmd.AddCommand(pauseCmd)

	pauseCmd.SetErrPrefix("Failed to pause session:")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pauseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pauseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}