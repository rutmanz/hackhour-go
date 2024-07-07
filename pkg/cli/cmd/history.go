package cmd

import (
	"fmt"
	"sort"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/rutmanz/hackhour-go/pkg/api"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:     "history",
	GroupID: "data",
	Short:   "Shows your HackHour session history",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := newClient()
		historyPtr, err := client.GetHistory()
		if err != nil {
			return err
		}

		history := *historyPtr

		if cmd.Flag("goal").Changed {
			goal := cmd.Flag("goal").Value.String()
			filteredHistory := []api.HistorySession{}
			for _, session := range history {
				if session.Goal == goal {
					filteredHistory = append(filteredHistory, session)
				}
			}
			history = filteredHistory
		}

		sort.Slice(history, func(i, j int) bool {
			if cmd.Flag("reverse").Changed {
				return history[i].CreatedAt.After(history[j].CreatedAt)
			}
			return history[i].CreatedAt.Before(history[j].CreatedAt)
		})
		tbl := table.NewWriter()
		tbl.SetAutoIndex(true)
		tbl.AppendHeader(table.Row{"Date", "Time", "Goal", "Activity"})
		total := 0
		for _, session := range history {
			tbl.AppendRow(table.Row{session.CreatedAt.Format("Mon Jan 1 15:05 PST"), session.Elapsed, session.Goal, session.Work})
			total += session.Elapsed
		}
		tbl.AppendFooter(table.Row{"", "", "Total", fmt.Sprintf("%d minutes", total)})
		fmt.Println(tbl.Render())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)

	historyCmd.Flags().BoolP("reverse", "r", false, "Sort with newest at the top")
	historyCmd.Flags().StringP("goal", "g", "", "Only show sessions from this goal")
}
