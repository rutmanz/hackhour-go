package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


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
}
