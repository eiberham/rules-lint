package linter

import (
	"regexp"
)

func jsToJSONString(content string) string {
	// Remove comments
	content = regexp.MustCompile(`/\*[\s\S]*?\*/`).ReplaceAllString(content, "")
	content = regexp.MustCompile(`(?m)//.*$`).ReplaceAllString(content, "")

	// Remove module.exports
	content = regexp.MustCompile(`module\.exports\s*=\s*`).ReplaceAllString(content, "")

	expression := regexp.MustCompile(`(\d+(?:\.\d+)?)\s*[*+\-/]\s*(\d+(?:\.\d+)?)(?:\s*[*+\-/]\s*(\d+(?:\.\d+)?))*`)

	// Find all mathematical expressions and convert them to strings
	content = expression.ReplaceAllStringFunc(content, func(match string) string {
		return `"` + match + `"`
	})

	// Wrap field names into double quotes
	content = regexp.MustCompile(`(\w+)\s*:`).ReplaceAllString(content, `"$1":`)

	// Transform single quotes to double quotes
	content = regexp.MustCompile(`'([^']*)'`).ReplaceAllString(content, `"$1"`)

	// Remove trailing commas before closing braces/brackets
	content = regexp.MustCompile(`,(\s*[}\]])`).ReplaceAllString(content, "$1")

	return content
}
