/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/rutmanz/hackhour-go/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hackhour",
	Short: "A hackhour api client",
	Long:  `A hackhour api client that allows you to interact with the hackhour api.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	viper.SetConfigName("config")             // name of config file (without extension)
	viper.SetConfigType("json")               // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.hackhour-go") // call multiple times to add many search paths
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found, creating one...")
			err := os.Mkdir(path.Join(os.Getenv("HOME"), ".hackhour-go"), 0755)
			if err != nil {
				fmt.Println(err)
			}

			viper.Set("api_key", os.Getenv("HACKHOUR_API_KEY"))
			viper.SafeWriteConfigAs(path.Join(os.Getenv("HOME"), ".hackhour-go", "config.json"))
		} else {
			// Config file was found but another error was produced
		}
	}
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func newClient() *api.HackHourClient {
	if viper.GetString("api_key") == "" {
		fmt.Println("Please login first")
		os.Exit(1)
	}
	return api.NewHackHourClient(viper.GetString("api_key"))
}
