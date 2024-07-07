/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// resumeCmd represents the resume command
var resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resumes the current session",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := newClient()
		session, err := client.GetSession()
		if err != nil {
			return err
		}
		if session.Paused {
			fmt.Println("Resuming session...")
			res, err := client.SessionPause()
			if err != nil {
				return err
			}
			if res.Paused {
				fmt.Println("Session could not be resumed")
			} else {
				fmt.Println("Session resumed")
			}
		} else {
			fmt.Println("Session already running")
		}
		return nil
	},
}

func init() {
	sessionCmd.AddCommand(resumeCmd)

	resumeCmd.SetErrPrefix("Failed to resume session:")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resumeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resumeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
