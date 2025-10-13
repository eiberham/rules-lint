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
	"bufio"
	"os"
)

type Handler interface {
	Next(h Handler) Handler
	Handle(rule Rule, cfg *Config, file *os.File) []string
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) Next(next Handler) Handler {
	h.next = next
	return next
}

func (h *BaseHandler) Handle(rule Rule, cfg *Config, file *os.File) []string {
	if h.next != nil {
		return h.next.Handle(rule, cfg, file)
	}
	return []string{}
}

type LineHandler struct {
	BaseHandler
}

type FileHandler struct {
	BaseHandler
}

func (h *LineHandler) Handle(rule Rule, cfg *Config, file *os.File) []string {
	issues := []string{}
	scanner := bufio.NewScanner(file)
	i := 0
	if rule.Type() == Line {
		for scanner.Scan() {
			i++
			line := scanner.Text()

			ctx := &RuleContext{
				Config:     cfg,
				Content:    line,
				LineNumber: i,
				FilePath:   file.Name(),
			}

			if issue := rule.Validate(ctx); issue != "" {
				issues = append(issues, issue)
			}
		}
		if len(issues) > 0 {
			return issues
		}
	}

	return h.BaseHandler.Handle(rule, cfg, file)
}

func (h *FileHandler) Handle(rule Rule, cfg *Config, file *os.File) []string {
	issues := []string{}
	if rule.Type() == File {
		content, err := os.ReadFile(file.Name())
		if err != nil {
			issues = append(issues, err.Error())
			return issues
		}
		ctx := &RuleContext{
			Config:   cfg,
			Content:  string(content),
			FilePath: file.Name(),
		}
		if issue := rule.Validate(ctx); issue != "" {
			issues = append(issues, issue)
		}
		if len(issues) > 0 {
			return issues
		}
	}

	return h.BaseHandler.Handle(rule, cfg, file)
}
