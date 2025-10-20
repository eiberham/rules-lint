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

type RuleType int

const (
	Line RuleType = iota
	File
)

type RuleContext struct {
	Config     *Config
	Content    string
	LineNumber int
	FilePath   string
}

type Rule interface {
	Validate(ctx *RuleContext) string
	Type() RuleType
}

type BaseRule struct {
	RuleType RuleType
}

func (r *BaseRule) Type() RuleType {
	return r.RuleType
}

type LineRule struct {
	BaseRule
	Handler func(line string, lineNumber int, ctx *RuleContext) string
}

func (r *LineRule) Validate(ctx *RuleContext) string {
	if r.Handler != nil {
		return r.Handler(ctx.Content, ctx.LineNumber, ctx)
	}
	return ""
}

type FileRule struct {
	BaseRule
	Handler func(content string, ctx *RuleContext) string
}

func (r *FileRule) Validate(ctx *RuleContext) string {
	if r.Handler != nil {
		return r.Handler(ctx.Content, ctx)
	}
	return ""
}

func isEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	left := make(map[string]bool)
	right := make(map[string]bool)

	for _, item := range a {
		left[item] = true
	}

	for _, item := range b {
		right[item] = true
	}

	for key := range left {
		if !right[key] {
			return false
		}
	}

	for key := range right {
		if !left[key] {
			return false
		}
	}

	return true
}
