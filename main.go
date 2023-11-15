/*
Copyright © 2023 Roger Ding
*/
package main

import (
	"github.com/roger-ding/command-line-todo-manager-go/sql"
	"github.com/roger-ding/command-line-todo-manager-go/cmd"
)

func main() {
  sql.Initialize()
  /*
  taskGet, err := sql.GetSingleTask(tasks, 1)
  if err != nil {
    panic(err)
  }
  sql.UpdateTaskStatus(*taskGet, "DONE")
  */
  cmd.Execute()
}
