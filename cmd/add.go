/*
Copyright © 2023 Roger Ding
*/
package cmd

import (
  "fmt"
  "strings"

  "github.com/roger-ding/command-line-todo-manager-go/sql"
  "github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
  Use:   "add",
  Short: "Adds a task to the todo task list",
  Example: `
  command-line-todo-manager-go add "task name"
  `,

  Run: func(cmd *cobra.Command, args []string) {
    taskToAdd := strings.Join(args, "")
    if len(taskToAdd) == 0 {
      fmt.Println("Task input is empty, please provide correct input!")
      return
    }

    sql.AddTask(taskToAdd, "incomplete")
    fmt.Printf("Added the following task to todo list: '%s'\n", taskToAdd)
  },
}

func init() {
  rootCmd.AddCommand(addCmd)
}
