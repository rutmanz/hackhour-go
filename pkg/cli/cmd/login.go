package cmd

import (
	"fmt"
	"os"

	"github.com/rutmanz/hackhour-go/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var loginCmd = &cobra.Command{
	Use:   "login [flags] [api key]",
	Args:  cobra.MaximumNArgs(1),
	Short: "Authorizes the client with your hackhour api key",
	Long: `Authorizes the client with your hackhour api key
	
	The api key can be found by running /api in the slack workspace`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var key string
		if v := os.Getenv("HACKHOUR_API_KEY"); v != "" {
			key = v
		}
		if len(args) >= 1 {
			key = args[0]
		}
		if key == "" {
			fmt.Println("Please provide an api key")
			cmd.Usage()
			return nil
		}
		fmt.Printf("Logging in...")

		client := api.NewHackHourClient(key)
		_, err := client.GetStats()
		if err != nil {
			return err
		}
		fmt.Print(" Success!\n")

		viper.Set("api_key", key)
		viper.WriteConfig()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.SetErrPrefix("Login failed:")
}
