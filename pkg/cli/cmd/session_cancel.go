package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
		printSimple(*session)
		return nil
	},
}

func init() {
	sessionCmd.AddCommand(cancelCmd)
	cancelCmd.SetErrPrefix("Failed to cancel session:")
}
