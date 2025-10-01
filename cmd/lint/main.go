package main

import (
	"fmt"
	"os"
	"path/filepath"
	"ruleslint/pkg/linter"
	"sync"
)

func main() {
	cfg, err := linter.LoadConfig("config.yaml")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	// Count .js files to determine buffer size
	var filePaths []string
	for _, dir := range cfg.Directories {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() && filepath.Ext(path) == ".js" {
				filePaths = append(filePaths, path)
			}
			return nil
		})
		if err != nil {
			fmt.Println("Error walking directory:", err)
		}
	}

	// Create channel with buffer size needed
	results := make(chan linter.LintResult, len(filePaths))
	var wg sync.WaitGroup
	linter.Init()

	// Spawn a goroutine for each file
	for _, filePath := range filePaths {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			linter.Run(path, cfg, results)
		}(filePath)
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
		}
	}
}
