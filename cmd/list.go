/*
Copyright © 2023 Roger Ding 
*/
package cmd

import (
  "fmt"

  "github.com/roger-ding/command-line-todo-manager-go/sql"
  "github.com/spf13/cobra"
)

const (
  incompleteBox = "\u2610"
  completeBox = "\u2611"
)

var (
  incompleteList []sql.TodoTask
  completedList []sql.TodoTask
  allList []sql.TodoTask
)

// listCmd represents the list command
var listCmd = &cobra.Command{
  Use:   "list",
  Short: "Lists all tasks from todo list, can specify incomplete and completed tasks",
  Example: `
  command-line-todo-manager-go list
    - lists all tasks
  command-line-todo-manager-go list -i
    - lists tasks that are incomplete
  command-line-todo-manager-go list -c
    - lists tasks that are completed
  `,

  Run: func(cmd *cobra.Command, args []string) {
    tasks := sql.GetAllTasks()
    for _, task := range tasks {
      if task.Status == "incomplete" {
        incompleteList = append(incompleteList, task)
      } 
      if task.Status == "completed" {
        completedList = append(completedList, task)
      }
      allList = append(allList, task)
    }

    var printList []sql.TodoTask
    if incompleteFlag, _ := cmd.Flags().GetBool("incomplete"); incompleteFlag {
      printList = incompleteList
    } else if completedFlag, _ := cmd.Flags().GetBool("completed"); completedFlag {
      printList = completedList
    } else {
      printList = allList 
    }

    for _, task := range printList {
      var boxStatus string
      if task.Status == "completed" {
        boxStatus = completeBox
      } else {
        boxStatus = incompleteBox
      }
      fmt.Println(boxStatus, task.Id, task.Task)
    }
  },
}

func init() {
  rootCmd.AddCommand(listCmd)
  listCmd.Flags().BoolP("incomplete", "i", false, "Lists incomplete todo tasks")
  listCmd.Flags().BoolP("completed", "c", false, "Lists completed todo tasks")
}
