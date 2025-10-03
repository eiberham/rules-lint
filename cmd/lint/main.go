package main

import (
	"fmt"
	"os"
	"ruleslint/pkg/linter"
	"sync"
)

func main() {
	cfg, err := linter.LoadConfig("config.yaml")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
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
		}
		if len(result.Issues) > 0 {
			fmt.Printf("Issues in %s:\n", result.FilePath)
			for _, issue := range result.Issues {
				fmt.Println(" -", issue)
			}
			os.Exit(1)
		}
	}

	fmt.Println("Linting completed successfully with no issues.")
	os.Exit(0)
}
