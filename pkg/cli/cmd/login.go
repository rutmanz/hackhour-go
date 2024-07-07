/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/rutmanz/hackhour-go/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login [flags] [api key]",
	Short: "Authorizes the client with your hackhour api key",
	Long: `Authorizes the client with your hackhour api key
	
	The api key can be found by running /api in the slack workspace`,
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		fmt.Printf("Logging in...")

		client := api.NewHackHourClient(key)
		_, err := client.GetStats()
		if err != nil {
			fmt.Printf("\nFailed to login: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Success!\n")
		
		viper.Set("api_key", key)
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	loginCmd.ValidArgs = []string{"APIKey"}
	loginCmd.Args = cobra.ExactArgs(1)

	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}