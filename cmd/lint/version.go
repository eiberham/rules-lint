package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  "All software has versions. This is the version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.0-dev")
	},
}
