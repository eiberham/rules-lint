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
