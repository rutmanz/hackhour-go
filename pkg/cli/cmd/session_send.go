package cmd

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/rutmanz/hackhour-go/pkg/cli"
	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Sends a message to your current HackHour session thread",
	RunE: func(cmd *cobra.Command, args []string) error {
		msg := strings.Join(args, " ")
		client := newSlackClient()

		var url string
		if cmd.Flag("git").Value.String() == "true" {
			url, err := cli.GetCommitURL()
			if err != nil {
				return err
			}
			msg = fmt.Sprintf("%v\n\n%v", msg, url)
		}

		if cmd.Flag("compare").Value.String() == "true" {
			session, err := client.GetHackHourClient().GetSession()
			if err != nil {
				return err
			}
			url, err = cli.GetCompareURL(&session.CreatedAt)
			if err != nil {
				return err
			}

			msg = fmt.Sprintf("%v\n\n%v", msg, url)
		}
		if msg == "" {
			return fmt.Errorf("message cannot be empty")
		}
		channel, ts, err := client.SendToSessionThread(msg, url)
		if err != nil {
			return err
		}
		fmt.Println("Message sent")
		if runtime.GOOS != "windows" {
			fmt.Println()
			if url != "" {
				fmt.Printf("\033]8;;%v\033\\Open Github Link\033]8;;\033\\\n", url)
			}
			fmt.Printf("\033]8;;https://hackclub.slack.com/archives/%v/p%v\033\\Open in Slack\033]8;;\033\\\n", channel, strings.Replace(ts, ".", "", 1))
		}
		return nil
	},
}

func init() {
	sessionCmd.AddCommand(sendCmd)
	sendCmd.SetErrPrefix("Failed to send to session:")
	sendCmd.Flags().BoolP("git", "g", false, "Include a link to the most recent git commit (supports github and gitlab)")
	sendCmd.Flags().BoolP("compare", "c", false, "Include a comparison between the most recent commit and HEAD from the start of the session (supports github)")
}
