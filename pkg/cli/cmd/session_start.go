package cmd

import (
	"fmt"
	"strings"

	"github.com/rutmanz/hackhour-go/pkg/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		fmt.Println("Session started")
		printSimple(*session)
		start_time, err := session.CreatedAt.MarshalText()
		if err != nil {
			return err
		}
		viper.Set("session_start", string(start_time))
		git_head, err := cli.GetGitHead()
		if err != nil {
			return err
		}
		viper.Set("session_start_commit", git_head)
		viper.WriteConfig()
		return nil
	},
}

func init() {
	sessionCmd.AddCommand(startCmd)
	startCmd.SetErrPrefix("Failed to start session:")
}
