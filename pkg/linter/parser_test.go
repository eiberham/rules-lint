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

import "testing"

func TestJsToJSONString(t *testing.T) {
	type TestSchema struct {
		name     string
		input    string
		expected string
	}

	tests := []TestSchema{
		{
			name:     "Simple string",
			input:    `{"key": "value"}`,
			expected: `{"key": "value"}`,
		},
		{
			name:     "String with single quotes",
			input:    `{'key': 'value'}`,
			expected: `{"key": "value"}`,
		},
		{
			name:     "String with mixed quotes",
			input:    `{"key": 'value'}`,
			expected: `{"key": "value"}`,
		},
		{
			name: "String with comments",
			input: `{
				// This is a comment
				"key": "value"
			}`,
			expected: `{
				
				"key": "value"
			}`,
		},
		{
			name:     "String with module.exports",
			input:    `module.exports = { "key": "value" }`,
			expected: `{ "key": "value" }`,
		},
		{
			name:     "String with mathematical expression",
			input:    `{"calculation": 5 + 3 * 2}`,
			expected: `{"calculation": "5 + 3 * 2"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := jsToJSONString(tt.input)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}
