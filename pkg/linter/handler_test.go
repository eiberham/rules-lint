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
	"os"
	"testing"
)

type TestLineHandler struct {
	BaseHandler
}

func (tlh *TestLineHandler) Handle(rule Rule, cfg *Config, file *os.File) []string {
	return []string{"Runs line handler"}
}

type TestFileHandler struct {
	BaseHandler
}

func (tfh *TestFileHandler) Handle(rule Rule, cfg *Config, file *os.File) []string {
	return []string{"Runs file handler"}
}

func TestHandler(t *testing.T) {

	rule := &LineRule{}

	cfg := &Config{}
	expected := []string{"Runs line handler"}

	processor := &BaseHandler{}
	processor.
		Next(&TestLineHandler{}).
		Next(&TestFileHandler{})

	received := processor.Handle(rule, cfg, nil)

	if !isEqual(received, expected) {
		t.Errorf("Expected %v, but got %v", expected, received)
	}

}
