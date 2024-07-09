package cmd

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Sends a message to your current HackHour session thread",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		msg := strings.Join(args, " ")
		client := newSlackClient()
		channel, ts, err := client.SendToSessionThread(msg)
		if err != nil {
			return err
		}
		fmt.Println("Message sent")
		if runtime.GOOS != "windows" {
			fmt.Printf("\033]8;;https://hackclub.slack.com/archives/%v/p%v\033\\Open in Slack\033]8;;\033\\\n", channel, strings.Replace(ts, ".", "", 1))
		}
		return nil
	},
}

func init() {
	sessionCmd.AddCommand(sendCmd)
	cancelCmd.SetErrPrefix("Failed to send to session:")
}
