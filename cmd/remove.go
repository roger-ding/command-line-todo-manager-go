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

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
  Use:   "remove",
  Short: "Removes a task from todo list based on id",
  Example: `
  command-line-todo-manager-go remove 1
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

    sql.DeleteTask(*taskGet)
    fmt.Printf("Removed the following task from todo list: '%s'\n", taskGet.Task)
  },
}

func init() {
  rootCmd.AddCommand(removeCmd)
}
