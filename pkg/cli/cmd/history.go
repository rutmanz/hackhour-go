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

		tbl := table.NewWriter()
		tbl.SetCaption("History")
		if cmd.Flag("goal").Changed {
			goal := cmd.Flag("goal").Value.String()
			filteredHistory := []api.HistorySession{}
			for _, session := range history {
				if session.Goal == goal {
					filteredHistory = append(filteredHistory, session)
				}
			}
			history = filteredHistory
			tbl.SetCaption("History for goal '%s'", goal)
		}

		sortBy := cmd.Flag("sort").Value.String()
		sort.Slice(history, func(i, j int) bool {
			switch sortBy {
			case "newest":
				return history[i].CreatedAt.Before(history[j].CreatedAt)
			case "oldest":
				return history[i].CreatedAt.After(history[j].CreatedAt)
			case "goal":
				if history[i].Goal == history[j].Goal {
					return history[i].CreatedAt.Before(history[j].CreatedAt)
				} else {
					return (history[i].Goal < history[j].Goal)
				}
			}
			return false
		})

		tbl.SetStyle(table.StyleRounded)
		tbl.SetAutoIndex(true)
		tbl.AppendHeader(table.Row{"Date", "Time", "Goal", "Activity"})
		total := 0
		lastGoal := ""
		for _, session := range history {
			if lastGoal != session.Goal && sortBy == "goal" {
				tbl.AppendSeparator()
				lastGoal = session.Goal
			}
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

	historyCmd.Flags().StringP("sort", "s", "newest", "Sorts history in by a given column. Options are: newest, oldest, goal")
	historyCmd.Flags().StringP("goal", "g", "", "Only show sessions from this goal")
}
