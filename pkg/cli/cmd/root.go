/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path"
	"reflect"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/rutmanz/hackhour-go/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "hackhour",
	Short: "A hackhour api client",
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
			os.Mkdir(path.Join(os.Getenv("HOME"), ".hackhour-go"), 0o755)

			viper.Set("api_key", os.Getenv("HACKHOUR_API_KEY"))
			viper.SafeWriteConfigAs(path.Join(os.Getenv("HOME"), ".hackhour-go", "config.json"))
		} else {
			// Config file was found but another error was produced
			fmt.Println("Error with config file")
		}
	}
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddGroup(&cobra.Group{
		ID:    "data",
		Title: "Queries:",
	})
}

func newClient() *api.HackHourClient {
	if viper.GetString("api_key") == "" {
		fmt.Println("Please login first")
		os.Exit(1)
	}
	return api.NewHackHourClient(viper.GetString("api_key"))
}

func printSimple[T any](data *T) {
	tbl := table.NewWriter()
	fields := reflect.VisibleFields(reflect.TypeOf(*data))
	tbl.SetStyle(table.StyleRounded)
	tbl.ResetHeaders()
	tbl.Style().Options.SeparateColumns = false
	tbl.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignRight},
		{Number: 2, Align: text.AlignLeft},
	})
	for _, field := range fields {
		tbl.AppendRow(table.Row{field.Name, reflect.ValueOf(*data).FieldByName(field.Name).Interface()})
	}

	fmt.Println(tbl.Render())
}
