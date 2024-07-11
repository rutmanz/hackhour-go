package cmd

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/rutmanz/hackhour-go/pkg/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			initial_time_str := viper.GetString("session_start")
			if initial_time_str == "" {
				return fmt.Errorf("no known session start time")
			}
			session, err := client.GetHackHourClient().GetSession()
			if err != nil {
				return err
			}
			initial_time := &time.Time{}
			err = initial_time.UnmarshalText([]byte(initial_time_str))
			if err != nil {
				return err
			}
			if (!session.CreatedAt.Equal(*initial_time)) {
				return fmt.Errorf("session start time does not match current session")
			}

			initial_commit := viper.GetString("session_start_commit")
			if initial_commit == "" {
				return fmt.Errorf("no known initial commit")
			}
			url, err = cli.GetCompareURL(initial_commit)
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
