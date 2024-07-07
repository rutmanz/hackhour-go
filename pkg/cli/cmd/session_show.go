package cmd

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)


var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows the state of the current session",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := newClient()
		session, err := client.GetSession()
		if err != nil {
			return err
		}
		if runtime.GOOS != "windows" {
			fmt.Printf("\033]8;;https://hackclub.slack.com/archives/C06SBHMQU8G/p%v\033\\Open in Slack\033]8;;\033\\\n\n", strings.Replace(session.MessageTs, ".", "", 1))
		}
		getJsonEncoder().Encode(session)
		return nil
	},
}

func init() {
	sessionCmd.AddCommand(showCmd)
}
