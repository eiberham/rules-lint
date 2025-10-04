package linter

import "testing"

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

func isEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
