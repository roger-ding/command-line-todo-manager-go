/*
Copyright © 2023 Roger Ding
*/
package main

import (
	"fmt"

	"github.com/roger-ding/command-line-todo-manager-go/sql"
	"github.com/roger-ding/command-line-todo-manager-go/cmd"
)

const (
  incompleteBox = "\u2610"
  completeBox = "\u2611"
)

func main() {
  sql.Initialize()
  sql.AddTask("wash dishes", "NOT DONE")
  sql.AddTask("take out trash", "NOT DONE")
  
  tasks := sql.GetAllTasks()
  for _, s := range tasks {
    var boxStatus string

    if s.Status == "DONE" {
      boxStatus = completeBox
    } else {
      boxStatus = incompleteBox
    }
    fmt.Println(boxStatus, s.Id, s.Task)
  }

  taskGet, err := sql.GetSingleTask(tasks, 1)
  if err != nil {
    panic(err)
  }
  sql.DeleteTask(*taskGet)
  sql.UpdateTaskStatus(*taskGet, "DONE")

  cmd.Execute()
}
