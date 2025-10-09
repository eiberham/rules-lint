package linter

import (
	"testing"
)

func TestLinter(t *testing.T) {
	const expected = "manual mock error"
	rule := &LineRule{
		BaseRule: BaseRule{RuleType: Line},
		Handler: func(line string, lineNumber int, ctx *RuleContext) string {
			return expected
		},
	}

	ctx := &RuleContext{}
	result := rule.Validate(ctx)

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}

	if rule.Type() != Line {
		t.Errorf("Expected Line type, got %v", rule.Type())
	}
}
