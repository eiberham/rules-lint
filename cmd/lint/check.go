/*
Copyright 2025 The Rules-Lint Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os"
	"ruleslint/pkg/linter"
	"sync"

	"github.com/spf13/cobra"
)

var check = &cobra.Command{
	Use:   "check",
	Short: "Check declarative rules files for issues",
	Long:  "Validate declarative rules files against configured rules ...",
	RunE:  run,
}

var config string

func init() {
	check.Flags().StringVar(&config, "config", "config.yml", "Path to configuration file")
}

func run(cmd *cobra.Command, args []string) error {
	cfg, err := linter.LoadConfig(config)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	paths := linter.GetFilesToLint(cfg)

	// Create channel with buffer size needed
	results := make(chan linter.LintResult, len(paths))
	var wg sync.WaitGroup
	linter.Init()

	// Spawn a goroutine for each file
	for _, path := range paths {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			linter.Run(path, cfg, results)
		}(path)
	}

	// Close channel when all goroutines are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Read results from the channel
	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error processing %s: %v\n", result.FilePath, result.Error)
			os.Exit(1)
		}
		if len(result.Issues) > 0 {
			fmt.Printf("Issues in %s:\n", result.FilePath)
			for _, issue := range result.Issues {
				fmt.Println(" -", issue)
			}
			os.Exit(1)
		}
	}

	fmt.Println("Linting completed with no issues.")
	return nil
}
