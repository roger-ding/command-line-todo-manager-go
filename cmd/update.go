/*
Copyright © 2023 Roger Ding 
*/
package cmd

import (
  "fmt"
  "strconv"
  "strings"

  "github.com/roger-ding/command-line-todo-manager-go/sql"
  "github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
  Use:   "update",
  Short: "Updates a task status",
  Example: `
  command-line-todo-manager-go update 1 -i
  command-line-todo-manager-go update 1 -c
  `,
  
  Run: func(cmd *cobra.Command, args []string) {
    convId, err := strconv.Atoi(strings.Join(args, ""))
    if err != nil {
        panic(err)
    }

    tasks := sql.GetAllTasks()
    taskGet, err := sql.GetSingleTask(tasks, convId)
    if err != nil {
      panic(err)
    }

    var stringComp string
    if incompleteFlag, _ := cmd.Flags().GetBool("incomplete"); incompleteFlag {
      stringComp = "incomplete"
    } else if completedFlag, _ := cmd.Flags().GetBool("completed"); completedFlag {
      stringComp = "completed"
    } else {
      panic("Flag not detected!\nMust be one of the following: [-i, -c]\nRefer to help method for more info!") 
    }
    sql.UpdateTaskStatus(*taskGet, stringComp)
    fmt.Printf("Updated the following task in todo list to %s status: '%s'\n", stringComp, taskGet.Task)
  },
}

func init() {
  rootCmd.AddCommand(updateCmd)
  updateCmd.Flags().BoolP("incomplete", "i", false, "Updates a task to incomplete status")
  updateCmd.Flags().BoolP("completed", "c", false, "Updates a task to completed status")
}
