package linter

import "testing"

func TestCheckTemplateVars(t *testing.T) {
	type TestSchema struct {
		name     string
		line     string
		expected string
	}

	ctx := &RuleContext{}

	tests := []TestSchema{
		{
			name:     "Valid template variable",
			line:     "'This is a valid ${template} variable'",
			expected: "",
		},
		{
			name:     "Unclosed template variable",
			line:     "'This is an unclosed ${template variable'",
			expected: "Line 2: Unclosed template variable: missing closing '}' in string",
		},
		{
			name:     "Orphaned closing brace",
			line:     "'This string has an orphaned } brace'",
			expected: "Line 3: Orphaned closing brace in string: missing opening '${' ",
		},
	}

	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CheckTemplateVars(test.line, i+1, ctx)
			if result != test.expected {
				t.Errorf("Test '%s' failed: expected '%s', got '%s'", test.name, test.expected, result)
			}
		})
	}
}

func TestCheckUnusedContextKeys(t *testing.T) {
	type TestSchema struct {
		name     string
		content  string
		expected string
	}

	ctx := &RuleContext{}

	test := TestSchema{
		name: "Unused context key",
		content: `{
			"context": {
				"foo" : {},
			}
		}`,
		expected: "Unused context key: foo",
	}

	t.Run(test.name, func(t *testing.T) {
		result := CheckUnusedContextKeys(test.content, ctx)
		if result != test.expected {
			t.Errorf("Test '%s' failed: expected '%s', got '%s'", test.name, test.expected, result)
		}
	})
}

func TestCheckAsyncIncongruence(t *testing.T) {
	type TestSchema struct {
		name     string
		content  string
		expected string
	}

	ctx := &RuleContext{}

	test := TestSchema{
		name: "Async incongruence",
		content: `{
			"async": true
		}`,
		expected: "Async rules should have matching conditions in the 'matches' section",
	}

	t.Run(test.name, func(t *testing.T) {
		result := CheckAsyncIncongruence(test.content, ctx)
		if result != test.expected {
			t.Errorf("Test '%s' failed: expected '%s', got '%s'", test.name, test.expected, result)
		}
	})

}
