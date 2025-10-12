package main

import (
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "rules-lint",
	Short: "A linter for declarative rule files",
	Long:  `A comprehensive linter that validates javascript rule files...`,
}

func init() {
	cmd.AddCommand(check)
}

func Run() error {
	return cmd.Execute()
}
