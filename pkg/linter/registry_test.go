package linter

import (
	"testing"
)

func TestRegistry(t *testing.T) {
	registry := NewRegistry()

	sampleRule := &LineRule{
		BaseRule: BaseRule{RuleType: Line},
		Handler: func(line string, lineNumber int, ctx *RuleContext) string {
			return ""
		},
	}

	registry.Register("sample-rule", sampleRule)

	retrievedRule, exists := registry.Get("sample-rule")
	if !exists {
		t.Fatalf("Expected rule 'sample-rule' to be registered")
	}

	if retrievedRule != sampleRule {
		t.Errorf("Retrieved rule does not match the registered rule")
	}

	_, exists = registry.Get("non-existent-rule")
	if exists {
		t.Errorf("Expected rule 'non-existent-rule' to not exist")
	}
}
