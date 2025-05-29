package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/thatquietkid/go-todo-app/internal/app"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [ID]",
	Short: "Delete a todo item and reassign IDs",
	Args:  cobra.ExactArgs(1),
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

		// Prepare new records and drop the deleted one
		newRecords := [][]string{header}
		found := false
		newID := 1

		for i, rec := range state.Records {
			if i == 0 {
				continue // skip header
			}

			id, err := strconv.Atoi(rec[colIndex["ID"]])
			if err != nil {
				continue
			}

			if id == taskID {
				found = true
				continue // skip this row (i.e., delete it)
			}

			// Update ID to maintain sequence
			rec[colIndex["ID"]] = strconv.Itoa(newID)
			newRecords = append(newRecords, rec)
			newID++
		}

		if !found {
			fmt.Println("Task ID not found:", taskID)
			return
		}

		state.Records = newRecords

		err = app.WriteCSVFile("tasks.csv", state.Records)
		if err != nil {
			fmt.Println("Error writing to CSV:", err)
			return
		}

		fmt.Printf("Task %d deleted and IDs updated successfully.\n", taskID)
		fmt.Println("Updated task list:")
		PrintTasks(state)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
