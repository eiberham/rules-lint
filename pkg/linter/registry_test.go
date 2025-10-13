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
