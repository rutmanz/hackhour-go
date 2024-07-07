package cmd

import (
	"github.com/spf13/cobra"
)

var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Manages your current HackHour session",
}

func init() {
	rootCmd.AddCommand(sessionCmd)
}
