package cmd

import (
	"fmt"
	"os"

	"github.com/rutmanz/hackhour-go/pkg/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var authslackCmd = &cobra.Command{
	Use:   "authslack [flags] [api key]",
	Args:  cobra.MaximumNArgs(1),
	Short: "Authorizes the client with a slack api key",
	Long: `Authorizes the client with a slack api key
	
	The api key needs the chat:write user scope`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var key string
		if v := os.Getenv("HACKHOUR_SLACK_TOKEN"); v != "" {
			key = v
		}
		if len(args) >= 1 {
			key = args[0]
		}
		if key == "" {
			fmt.Println("Please provide a slack api key")
			cmd.Usage()
			return nil
		}
		fmt.Printf("Logging in...")

		client := slack.CreateClient(newClient(), key)
		identity, err := client.CheckAuth()
		if err != nil {
			return err
		}
		fmt.Printf(" Success!\nAuthenticated as %v\n", identity.User)
		viper.Set("slack_user", identity.UserID)
		viper.Set("slack_token", key)
		viper.WriteConfig()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(authslackCmd)
	authslackCmd.SetErrPrefix("Slack login failed:")
}
