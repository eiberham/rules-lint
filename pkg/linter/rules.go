package linter

import (
	"fmt"
	"regexp"
	"strings"
)

type RuleContext struct {
	Config *Config
	// Could add other context data in the future like:
	// FileContent string  // for full-file rules
}

type Rule func(line string, lineNumber int, context *RuleContext) string

var Rules = map[string]Rule{
	"checkTemplateVars":     checkTemplateVars,
	"checkValidActionTypes": checkValidActionTypes,
}

func checkTemplateVars(line string, lineNumber int, ctx *RuleContext) string {
	var issues []string

	// Only check template variables within string contexts (single quotes)
	// Find all single-quoted strings first
	pattern := regexp.MustCompile(`'([^']*)'`)
	matches := pattern.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		if len(match) > 1 {
			content := match[1] // The content inside the quotes

			// 1. Check for unclosed ${ within the string
			// \$\{[^}]*$ breakdown:
			// \$ - matches literal $
			// \{ - matches literal {
			// [^}]* - matches any characters except } (zero or more times)
			// $ - asserts position at end of string
			if regexp.MustCompile(`\$\{[^}]*$`).MatchString(content) {
				issues = append(issues, "Unclosed template variable: missing closing '}' in string")
				continue
			}

			// 2. Check for orphaned } within the string (} without ${)
			// Breakdown:
			// \} - matches literal }
			// \$\{ - matches literal ${
			if regexp.MustCompile(`\}`).MatchString(content) && !regexp.MustCompile(`\$\{`).MatchString(content) {
				issues = append(issues, "Orphaned closing brace in string: missing opening '${' ")
				continue
			}

			// 3. Check for plain {...} without $ prefix within the string
			// \{[^}]*\} breakdown:
			// \{ - matches literal {
			// [^}]* - matches any characters except } (zero or more times)
			// \} - matches literal }
			if regexp.MustCompile(`\{[^}]*\}`).MatchString(content) && !regexp.MustCompile(`\$\{`).MatchString(content) {
				issues = append(issues, "Found '{...}' in string - should be '${...}' for template variables")
				continue
			}
		}
	}

	if len(issues) > 0 {
		return fmt.Sprintf("Line %d: %s", lineNumber, strings.Join(issues, "; "))
	}

	return ""
}

func checkValidActionTypes(line string, lineNumber int, ctx *RuleContext) string {
	// Define valid action types
	return ""
}
