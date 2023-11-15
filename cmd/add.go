/*
Copyright © 2023 Roger Ding
*/
package cmd

import (
  "fmt"

  //"github.com/roger-ding/command-line-todo-manager-go/sql"
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
    fmt.Println("add called")
    fmt.Println(args[0])
  },
}

func init() {
  rootCmd.AddCommand(addCmd)
}
