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
	RunE: func(cmd *cobra.Command, args []string) error {
		client := newClient()
		session, err := client.SessionCancel()
		if err != nil {
			return err
		}
		fmt.Println("Session cancelled:", session.ID)
		getJsonEncoder().Encode(session)
		return nil
	},
}

func init() {
	sessionCmd.AddCommand(cancelCmd)
	cancelCmd.SetErrPrefix("Failed to cancel session:")	
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cancelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cancelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
