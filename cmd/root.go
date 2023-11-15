/*
Copyright © 2023 Roger Ding
*/
package cmd

import (
  "os"

  "github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "command-line-todo-manager-go",
  Short: "\n*** Command line todo list manager ***",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  err := rootCmd.Execute()
  if err != nil {
    os.Exit(1)
  }
}

func init() {
  // Disable auto completion feature
  rootCmd.CompletionOptions.DisableDefaultCmd = true 
}
