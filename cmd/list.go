package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
	"strings"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
	"github.com/thatquietkid/go-todo-app/internal/app"
)

func PrintTasks(state *app.AppState) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)
	for i, rec := range state.Records {
		if i == 0 {
			fmt.Fprintf(writer, "%s\t%s\n", strings.Join(rec, "\t"), "Time Left")
			continue
		}

		dueTime, err := time.Parse("2006-01-02", rec[3])
		relativeDue := "invalid date"
		if err == nil {
			relativeDue = timediff.TimeDiff(dueTime)
		}

		fmt.Fprintf(writer, "%s\t%s\n", strings.Join(rec, "\t"), relativeDue)
	}
	writer.Flush()
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		state, ok := ctx.Value("appState").(*app.AppState)
		if !ok || state == nil {
			fmt.Println("App state is not available")
			return
		}

		PrintTasks(state)
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
