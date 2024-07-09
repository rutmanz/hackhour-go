package cmd

import (
	"fmt"
	"path"
	"regexp"
	"runtime"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Sends a message to your current HackHour session thread",
	RunE: func(cmd *cobra.Command, args []string) error {
		msg := strings.Join(args, " ")
		client := newSlackClient()
		if cmd.Flag("git").Value.String() == "true" {
			repo, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{DetectDotGit: true})
			if err != nil {
				return err
			}
			remotes, err := repo.Remotes()
			if err != nil {
				return err
			}
			if len(remotes) == 0 {
				return fmt.Errorf("no remotes found")
			}
			remote := remotes[0]
			git_origin := strings.TrimSpace(remote.Config().URLs[0])
			git_origin = regexp.MustCompile(`\.git$`).ReplaceAllString(git_origin, "")                           // remove trailing .git
			git_origin = regexp.MustCompile(`\w+(:\w+)?@github\.com`).ReplaceAllString(git_origin, "github.com") // remove usernames and tokens from url
			ref, err := repo.Reference("HEAD", true)
			if err != nil {
				return err
			}
			latest_commit := ref.Hash().String()
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
	sendCmd.Flags().BoolP("git", "g", false, "Include a link to the most recent git commit (supports github and gitlab)")
}
