package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

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
}
