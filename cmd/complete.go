package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/thatquietkid/go-todo-app/internal/app"
)

var completeCmd = &cobra.Command{
	Use:     "complete [taskID]",
	Aliases: []string{"done"},
	Short:   "Mark a task as completed",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		state, ok := ctx.Value("appState").(*app.AppState)
		if !ok || state == nil {
			fmt.Println("App state is not available")
			return
		}

		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID:", args[0])
			return
		}

		header := state.Records[0]
		colIndex := make(map[string]int)
		for i, col := range header {
			colIndex[col] = i
		}

		found := false
		for i, rec := range state.Records {
			if i == 0 {
				continue
			}

			id, err := strconv.Atoi(rec[colIndex["ID"]])
			if err != nil {
				continue
			}
			if id == taskID {
				state.Records[i][colIndex["Completed"]] = "true"
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Task ID not found:", taskID)
			return
		}

		err = app.WriteCSVFile("tasks.csv", state.Records)
		if err != nil {
			fmt.Println("Error writing to CSV:", err)
			return
		}
		fmt.Println("Task", taskID, "marked as completed.")
		fmt.Println("Updated task list:")
		// Print the updated task list
		PrintTasks(state)
	},
}

func init() {
	RootCmd.AddCommand(completeCmd)
}
