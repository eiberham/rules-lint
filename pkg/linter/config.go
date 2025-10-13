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
	"path/filepath"

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

func GetFilesToLint(cfg *Config) []string {
	var paths []string
	for _, dir := range cfg.Directories {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && filepath.Ext(path) == ".js" {
				paths = append(paths, path)
			}

			return nil
		})

		if err != nil {
			fmt.Printf("Error walking directory %s: %v\n", dir, err)
			return nil
		}
	}

	return paths
}
