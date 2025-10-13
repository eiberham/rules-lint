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
