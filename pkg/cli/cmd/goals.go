package cmd

import (
	"fmt"
	"sort"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var goalsCmd = &cobra.Command{
	Use:     "goals",
	GroupID: "data",
	Short:   "Shows your HackHour goals list",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := newClient()
		goalsPtr, err := client.GetGoals()
		if err != nil {
			return err
		}
		goals := *goalsPtr
		sort.Slice(goals, func(i, j int) bool {
			if cmd.Flag("desc").Changed {
				return goals[i].Minutes > goals[j].Minutes
			}
			return goals[i].Minutes < goals[j].Minutes
		})
		tbl := table.NewWriter()
		tbl.SetStyle(table.StyleRounded)
		tbl.SetAutoIndex(true)
		tbl.AppendHeader(table.Row{"Goal", "Minutes"})
		for _, goal := range goals {
			tbl.AppendRow(table.Row{goal.Name, goal.Minutes})
		}
		fmt.Println(tbl.Render())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(goalsCmd)

	goalsCmd.Flags().BoolP("desc", "d", false, "Sorts goals in descending order")
}
