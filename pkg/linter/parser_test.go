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
