package linter

import (
	"bufio"
	"log"
	"os"
)

type LintResult struct {
	FilePath string
	Issues   []string
	Error    error
}

func Run(path string, cfg *Config, results chan<- LintResult) {
	// Placeholder for linting logic
	// This function would read the file, apply the rules from cfg,
	// and report any issues found.

	file, err := os.Open(path)
	if err != nil {
		results <- LintResult{FilePath: path, Error: err}
		return
	}
	defer file.Close()

	issues := []string{}
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	// Iterate through the lines of the file
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		ctx := &RuleContext{Config: cfg}
		for ruleName, rule := range Rules {
			// Check if this rule is enabled in config
			if cfg.Rules[ruleName] {
				if issue := rule(line, lineNumber, ctx); issue != "" {
					issues = append(issues, issue)
				}
			}
		}
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	results <- LintResult{FilePath: path, Issues: issues}
}
