package cmd

import (
	"fmt"
	"strconv"
	"time"
	"strings"
	"github.com/spf13/cobra"
	"github.com/thatquietkid/go-todo-app/internal/app"
)

var addCmd = &cobra.Command{
	Use:   "add [title] [description] [due-date] [completed]",
	Short: "Add a new todo item",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		state, ok := ctx.Value("appState").(*app.AppState)
		if !ok || state == nil {
			fmt.Println("App state is not available")
			return
		}

		title := args[0]
		description := args[1]
		dueDateStr := args[2]
		completedStr := args[3]
		
		// Validate due date format
		dueDateStr = strings.TrimSpace(dueDateStr)
		dueDate, err := time.Parse("2006-01-02", dueDateStr)
		if err != nil {
			fmt.Printf("Invalid due date format: '%s'. Expected YYYY-MM-DD (e.g. 2025-05-30).\n", dueDateStr)
			return
		}


		// Validate completed value
		_, err = strconv.ParseBool(completedStr)
		if err != nil {
			fmt.Println("Invalid completed value. Expected true or false.")
			return
		}

		// Get the header row and column index map
		header := state.Records[0]
		colIndex := make(map[string]int)
		for i, col := range header {
			colIndex[col] = i
		}

		// Generate new ID
		newID := 1
		for i, row := range state.Records {
			if i == 0 {
				continue
			}
			id, err := strconv.Atoi(row[colIndex["ID"]])
			if err == nil && id >= newID {
				newID = id + 1
			}
		}

		// Create new row
		newRow := make([]string, len(header))
		for i, col := range header {
			switch col {
			case "ID":
				newRow[i] = strconv.Itoa(newID)
			case "Title":
				newRow[i] = title
			case "Description":
				newRow[i] = description
			case "Due Date":
				newRow[i] = dueDate.Format("2006-01-02")
			case "Completed":
				newRow[i] = completedStr
			default:
				newRow[i] = "" // fill any unused column
			}
		}

		// Append new row and save
		state.Records = append(state.Records, newRow)

		err = app.WriteCSVFile("tasks.csv", state.Records)
		if err != nil {
			fmt.Println("Error writing to CSV:", err)
			return
		}

		fmt.Printf("Todo added with ID %d.\n", newID)
		fmt.Println("Updated task list:")
		PrintTasks(state)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
