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
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	file, err := os.CreateTemp("", "config_test_*.yml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())
	defer file.Close()

	content := strings.TrimSpace(`
directories:
  - ./rules
`)

	if _, err := file.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	cfg, err := LoadConfig(file.Name())
	fmt.Printf("%+v\n", cfg)

	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if len(cfg.Directories) != 1 || cfg.Directories[0] != "./rules" {
		t.Errorf("Expected directories to be ['./rules'], got %v", cfg.Directories)
	}

}
func TestGetEnabledRules(t *testing.T) {
	type TestSchema struct {
		name     string
		config   *Config
		expected []string
	}

	tests := []TestSchema{
		{
			name: "No rules enabled",
			config: &Config{
				Rules: map[string]bool{},
			},
			expected: []string{},
		},
		{
			name: "Some rules enabled",
			config: &Config{
				Rules: map[string]bool{
					"rule1": true,
					"rule2": false,
					"rule3": true,
				},
			},
			expected: []string{"rule1", "rule3"},
		},
		{
			name: "All rules enabled",
			config: &Config{
				Rules: map[string]bool{
					"rule1": true,
					"rule2": true,
					"rule3": true,
				},
			},
			expected: []string{"rule1", "rule2", "rule3"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GetEnabledRules(test.config)
			if !isEqual(result, test.expected) {
				t.Errorf("Test '%s' failed: expected %v, got %v", test.name, test.expected, result)
			}
		})
	}
}
