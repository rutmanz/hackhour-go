package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"path"
	"regexp"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Sends a message to your current HackHour session thread",
	RunE: func(cmd *cobra.Command, args []string) error {
		msg := strings.Join(args, " ")
		client := newSlackClient()
		if cmd.Flag("git").Value.String() == "true" {
			git_origin_cmd := exec.Command("git", "config", "--get", "remote.origin.url")
			var out bytes.Buffer
			git_origin_cmd.Stdout = &out
			err := git_origin_cmd.Run()
			if err != nil {
				return err
			}
			git_origin := strings.TrimSpace(out.String())
			regexp.MustCompile(`/\.git^/`).ReplaceAllString(git_origin, "")
			regexp.MustCompile(`/:\w+@/`).ReplaceAllString(git_origin, "")
			latest_commit_cmd := exec.Command("git", "rev-parse", "HEAD")
			out.Reset()
			latest_commit_cmd.Stdout = &out
			err = latest_commit_cmd.Run()
			if err != nil {
				return err
			}
			latest_commit := strings.TrimSpace(out.String())
			url := strings.Replace(path.Join(git_origin, "commit", latest_commit), ":/", "://", 1) // path.Join normalizes https:// to http:/
			msg = fmt.Sprintf("%v\n%v", msg, url)
		}
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
	sendCmd.SetErrPrefix("Failed to send to session:")
	sendCmd.Flags().BoolP("git", "g", false, "Include a link to the most recent git commit (uses remote 'origin'; supports github)")
}
