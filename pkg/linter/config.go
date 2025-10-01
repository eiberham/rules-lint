package linter

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Directories []string        `yaml:"directories"`
	Rules       map[string]bool `yaml:"rules"`
}

// TODO:
// Add support for JSON config files
// Add support for JS config files (parsing the object from module.exports)
// Add support for .ruleslintrc files (parsing the JSON object)
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func GetEnabledRules(cfg *Config) []string {
	var enabled []string
	for rule, ok := range cfg.Rules {
		if ok {
			enabled = append(enabled, rule)
		}
	}
	return enabled
}
