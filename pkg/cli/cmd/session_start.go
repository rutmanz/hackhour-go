/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start [message...]",
	Short: "Starts a new arcade session",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		msg := strings.Join(args, " ")
		client := newClient()
		session, err := client.SessionStart(msg)
		if err != nil {
			return err
		}
		fmt.Println("Session started:", session.ID)
		return nil
	},
}

func init() {
	sessionCmd.AddCommand(startCmd)
	startCmd.SetErrPrefix("Failed to start session:")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
