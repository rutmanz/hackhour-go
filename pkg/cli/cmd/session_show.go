/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
