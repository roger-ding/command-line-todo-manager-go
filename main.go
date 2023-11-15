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
  cmd.Execute()
}
