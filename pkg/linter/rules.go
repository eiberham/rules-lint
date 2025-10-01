package linter

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func CheckTemplateVars(line string, lineNumber int, ctx *RuleContext) string {
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

func CheckUnusedContextKeys(content string, ctx *RuleContext) string {
	str := jsToJSONString(content)

	var result map[string]interface{}
	err := json.Unmarshal([]byte(str), &result)
	if err != nil {
		return fmt.Sprintf("Error parsing JSON: %v", err)
	}

	context, ok := result["context"].(map[string]interface{})
	if !ok {
		return "No context section found or context is not an object"
	}

	var contextKeys []string
	for key := range context {
		contextKeys = append(contextKeys, key)
	}

	var issues []string
	for _, val := range contextKeys {
		if strings.Count(str, val+"/") == 0 {
			issues = append(issues, fmt.Sprintf("Unused context key: %s", val))
		}
	}

	if len(issues) > 0 {
		return strings.Join(issues, "\n")
	}

	return ""
}

func CheckAsyncIncongruence(content string, ctx *RuleContext) string {
	str := jsToJSONString(content)

	var result map[string]interface{}
	err := json.Unmarshal([]byte(str), &result)
	if err != nil {
		return fmt.Sprintf("Error parsing JSON: %v", err)
	}

	async, ok := result["async"].(bool)
	if !ok {
		async = false // Default to false if not present or not a boolean
	}

	// Check if matches section exists and has content
	consistent := false
	if matches, ok := result["matches"].(map[string]interface{}); ok {
		// Check if matches has any content
		if len(matches) > 0 {
			// Check if "all" or "any" arrays have content
			if all, exists := matches["all"].([]interface{}); exists && len(all) > 0 {
				consistent = true
			} else if any, exists := matches["any"].([]interface{}); exists && len(any) > 0 {
				consistent = true
			}
		}
	}

	if async && !consistent {
		return "Async rules should have matching conditions in the 'matches' section"
	}

	return ""
}
